package main

import (
	. "github.com/dave/jennifer/jen"
)

func writeCtor(f *File, structName string, entInfo *entStructInfo, fields []*fieldInfo) {

	// 写 EmptyXXX
	f.Func().Id(EmptyCtor(structName)).Params().Op("*").Id(structName).
		Block(
			Return(Id(NormalCtor(structName)).CallFunc(func(g *Group) {
				for _, field := range fields {
					g.Add(field.emptyValue)
				}
			})),
		)
	// 写 NewXXX
	f.Func().Id(NormalCtor(structName)).ParamsFunc(func(g *Group) {
		for _, field := range fields {
			g.Add(field.setParam)
		}
	}).Op("*").Id(structName).
		BlockFunc(func(g *Group) {

			g.Id("m").Op(":=").Parens(Op("*").Id(structName)).Parens(Qual(attrPackageName, "NewStrMap").Call(Nil()))

			for _, field := range fields {
				g.Id("m").Dot("").Add(field.setter).Call(Id(field.key))
			}

			g.Id("m").Dot("ClearChangeKey").Call()

			if entInfo != nil {
				// 默认的属性，位置，朝向和 ID
				// m.SetPos(attr.EmptyVec3())
				// m.SetRot(attr.EmptyVec3())
				// m.SetID("")
				g.Comment("实体内置的属性")

				g.Comment("实体内置的 ID")
				g.Id("m").Dot("SetId").Call(Lit(""))

				if entInfo.pos != nil {
					g.Comment("实体内置的 位置")
					g.Id("m").Dot("SetPos").Call(Qual(attrPackageName, "EmptyVec3").Call())
				}
				if entInfo.rot != nil {
					g.Comment("实体内置的 朝向")
					g.Id("m").Dot("SetRot").Call(Qual(attrPackageName, "EmptyVec3").Call())
				}
			}

			g.Return(Id("m"))
		})
}
