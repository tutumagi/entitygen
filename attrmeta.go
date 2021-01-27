package main

import (
	"go/types"
	"strings"

	. "github.com/dave/jennifer/jen"
)

func writeAttrDef(f *File, attrDefName string, fields []*structField) {
	// *attr.Def
	attrDef := func() *Statement { return Id("*").Qual("entitygen/attr", "Def") }

	// var xxxAttrDef *attr.Def
	f.Var().Id(attrDefName).Add(attrDef())
	f.Func().Id("init").Params().
		BlockFunc(
			func(g *Group) {
				g.Id(attrDefName).Op("=").Op("&").Qual("entitygen/attr", "Def").Block()
				g.Line()

				for i := 0; i < len(fields); i++ {
					field := fields[i]

					switch v := field.typ.(type) {
					case *types.Basic:
						g.Id(attrDefName).Dot("DefAttr").CallFunc(func(ig *Group) {
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
