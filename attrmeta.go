package main

import (
	"go/types"
	"strings"

	. "github.com/dave/jennifer/jen"
)

func writeAttrMeta(f *File, attrMetaName string, structName string, fields []*structField) {
	// var xxxAttrDef *attr.Def
	f.Var().Id(attrMetaName).Id("*").Add(attrMeta())
	f.Func().Id("init").Params().
		BlockFunc(
			func(g *Group) {

				// attr.NewMeta(func() interface{} {
				// 	return EmptyDesk()
				// }, func() interface{} {
				// 	return []*Desk{}
				// })

				g.Id(attrMetaName).Op("=").Add(attrNewMeta()).Call(Func().Params().Interface().Block(
					Return(Id(EmptyCtor(structName)).Call()),
				), Func().Params().Interface().Block(
					Return(Op("&").Index().Op("*").Id(structName).Block()),
				))
				g.Line()

				for i := 0; i < len(fields); i++ {
					field := fields[i]

					g.Id(attrMetaName).Dot("DefAttr").CallFunc(func(ig *Group) {
						ig.Lit(field.key)
						switch v := field.typ.(type) {
						case *types.Basic:
							ig.Qual(attrPackageName, strings.Title(v.Name()))
						default:
							ig.Op("&").Id(trimHeadStar(getTypString(v))).Block()
						}

						if field.cell {
							ig.Qual(attrPackageName, "AfOtherClients")
						} else {
							ig.Qual(attrPackageName, "AfBase")
						}

						if field.storeDB {
							ig.True()
						} else {
							ig.False()
						}
					})

				}

			},
		)
}
