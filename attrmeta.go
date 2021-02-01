package main

import (
	"go/types"
	"strings"

	. "github.com/dave/jennifer/jen"
)

func writeAttrMeta(f *File, attrMetaName string, fields []*structField) {
	// *attr.Def
	attrMeta := func() *Statement { return Id("*").Qual(attrPackageName, "Meta") }

	// var xxxAttrDef *attr.Def
	f.Var().Id(attrMetaName).Add(attrMeta())
	f.Func().Id("init").Params().
		BlockFunc(
			func(g *Group) {
				g.Id(attrMetaName).Op("=").Op("&").Qual(attrPackageName, "Meta").Block()
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
							ig.Qual(attrPackageName, "AfCell")
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
