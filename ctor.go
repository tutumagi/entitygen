package main

import (
	"fmt"

	. "github.com/dave/jennifer/jen"
)

func writeCtor(f *File, structName string, sourceTypeName string, fields []*structField) {
	// EmptyXXXX 和 NewXXX
	emptyCtorName := fmt.Sprintf("Empty%s", sourceTypeName)
	normalCtorName := fmt.Sprintf("New%s", sourceTypeName)
	// 写 EmptyXXX
	f.Func().Id(emptyCtorName).Params().Op("*").Id(structName).
		Block(
			Return(Id(normalCtorName).CallFunc(func(g *Group) {
				for _, field := range fields {
					g.Add(field.emptyValue)
				}
			})),
		)
	// 写 NewXXX
	f.Func().Id(normalCtorName).ParamsFunc(func(g *Group) {
		for _, field := range fields {
			g.Add(field.setParam)
		}
	}).Op("*").Id(structName).
		BlockFunc(func(g *Group) {

			g.Id("m").Op(":=").Parens(Op("*").Id(structName)).Parens(Qual("entitygen/attr", "NewStrMap").Call(Nil()))

			for _, field := range fields {
				g.Id("m").Dot("").Add(field.setter).Call(Id(field.key))
			}

			g.Id("m").Dot("ClearChangeKey").Call()

			g.Return(Id("m"))
		})
}
