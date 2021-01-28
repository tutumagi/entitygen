package main

import (
	"go/types"
	"strings"

	. "github.com/dave/jennifer/jen"
)

func writeAttrMeta(f *File, attrMetaName string, fields []*structField) {
	// *attr.Def
	attrMeta := func() *Statement { return Id("*").Qual("entitygen/attr", "Meta") }

	// var xxxAttrDef *attr.Def
	f.Var().Id(attrMetaName).Add(attrMeta())
	f.Func().Id("init").Params().
		BlockFunc(
			func(g *Group) {
				g.Id(attrMetaName).Op("=").Op("&").Qual("entitygen/attr", "Meta").Block()
				g.Line()

				for i := 0; i < len(fields); i++ {
					field := fields[i]

					switch v := field.typ.(type) {
					case *types.Basic:
						g.Id(attrMetaName).Dot("DefAttr").CallFunc(func(ig *Group) {
							ig.Lit(field.key)
							ig.Qual("entitygen/attr", strings.Title(v.Name()))

							if field.cell {
								ig.Qual("entitygen/attr", "AfCell")
							} else {
								ig.Qual("entitygen/attr", "AfBase")
							}

							if field.storeDB {
								ig.True()
							} else {
								ig.False()
							}
						})
					}
				}
			},
		)
}
