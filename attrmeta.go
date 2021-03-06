package main

import (
	"go/types"
	"strings"

	. "github.com/dave/jennifer/jen"
)

func writeAttrMeta(f *File, attrMetaName string, structName string, entInfo *entStructInfo, fields []*fieldInfo) {
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

				// 写自定义属性
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

				// 写实体内置的属性定义
				if entInfo != nil {
					// 默认的属性，位置，朝向和 ID
					// RoomDefMeta.DefAttr("id", attr.String, attr.AfOtherClients, true)
					// RoomDefMeta.DefAttr("pos", attr.Vector3, attr.AfBase, true)
					// RoomDefMeta.DefAttr("rot", attr.Vector3, attr.AfBase, true)

					g.Comment("实体内置的属性")

					g.Comment("实体内置的 ID")
					g.Id(attrMetaName).Dot("DefAttr").Call(
						Lit(buildinIDKey),
						Qual(attrPackageName, "String"),
						Qual(attrPackageName, "AfOtherClients"),
						True(), // id 肯定需要写到 db 里面去
					)

					if entInfo.pos != nil {
						g.Comment("实体内置的 位置")
						g.Id(attrMetaName).Dot("DefAttr").CallFunc(
							func(ig *Group) {
								ig.Lit(buildinPosKey)
								ig.Qual(attrPackageName, "Vector3")
								if entInfo.pos.hasCell {
									ig.Qual(attrPackageName, "AfOtherClients")
								} else {
									ig.Qual(attrPackageName, "AfBase")
								}

								if entInfo.pos.storedb {
									ig.True()
								} else {
									ig.False()
								}
							},
						)
					}
					if entInfo.rot != nil {
						g.Comment("实体内置的 朝向")
						g.Id(attrMetaName).Dot("DefAttr").CallFunc(
							func(ig *Group) {
								ig.Lit(buildinRotKey)
								ig.Qual(attrPackageName, "Vector3")
								if entInfo.rot.hasCell {
									ig.Qual(attrPackageName, "AfOtherClients")
								} else {
									ig.Qual(attrPackageName, "AfBase")
								}

								if entInfo.rot.storedb {
									ig.True()
								} else {
									ig.False()
								}
							},
						)
					}
				}
			},
		)
}
