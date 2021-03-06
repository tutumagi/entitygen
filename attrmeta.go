package main

import (
	"go/types"
	"strings"

	. "github.com/dave/jennifer/jen"
)

func writeAttrMeta(f *File, attrMetaName string, structName string, isEntity bool, fields []*fieldInfo) {
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
				if isEntity {
					// 默认的属性，位置，朝向和 ID
					// RoomDefMeta.DefAttr("id", attr.String, attr.AfOtherClients, true)
					// RoomDefMeta.DefAttr("pos", attr.Vector3, attr.AfBase, true)
					// RoomDefMeta.DefAttr("rot", attr.Vector3, attr.AfBase, true)

					if EntityGenID || EntityGenPos || EntityGenRot {
						g.Comment("实体内置的属性")
					}
					if EntityGenID {
						g.Comment("实体内置的 ID")
						g.Id(attrMetaName).Dot("DefAttr").Call(
							Lit(buildinIDKey),
							Qual(attrPackageName, "String"),
							Qual(attrPackageName, "AfOtherClients"),
							True(), // id 肯定需要写到 db 里面去
						)
					}
					if EntityGenPos {
						g.Comment("实体内置的 位置")
						g.Id(attrMetaName).Dot("DefAttr").Call(
							Lit(buildinPosKey),
							Qual(attrPackageName, "Vector3"),
							Qual(attrPackageName, "AfCell"),
							True(), // TODO 位置需不需要写到 db 里面去，应该来外部来配置，不是每种实体都需要记录位置到 db
						)
					}
					if EntityGenRot {
						g.Comment("实体内置的 朝向")
						g.Id(attrMetaName).Dot("DefAttr").Call(
							Lit("rot"),
							Qual(attrPackageName, "Vector3"),
							Qual(attrPackageName, "AfCell"),
							True(), // TODO 朝向需不需要写到 db 里面去，应该来外部来配置，不是每种实体都需要记录位置到 db
						)
					}
				}
			},
		)
}
