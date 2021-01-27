package main

import (
	"fmt"
	"go/types"

	. "github.com/dave/jennifer/jen"
)

func writeGetterSetter(f *File, fields []*structField, thisFn func() *Statement, convertThisFn func() *Statement) error {
	for i := 0; i < len(fields); i++ {

		field := fields[i]

		switch v := field.typ.(type) {
		case *types.Basic, *types.Pointer:
			// 写 getter
			_, isBasic := v.(*types.Basic)
			// func (a *XXXDef) GetField() FieldType
			f.Func().Params(thisFn()).Add(field.getter).Params().Id(field.typString).
				Block(
					ReturnFunc(func(g *Group) {
						statement := g.Add(convertThisFn()).Dot(field.attrGetter).Params(Lit(field.key))
						// 如果不是基础类型，则加上类型转换
						if !isBasic {
							statement.Dot("").Parens(Id(field.typString))
						}
					}),
				)

			//  写 setter
			f.Func().Params(thisFn()).Add(field.setter).Params(field.setParam).
				Block(
					convertThisFn().Dot("Set").Params(Lit(field.key), Id(field.key)),
				)

			// 换行符
			f.Line()
		case *types.Map:
			switch mapK := v.Key().(type) {
			case *types.Basic:
				if mapK.Kind() == types.Int32 || mapK.Kind() == types.String {

				} else {
					return fmt.Errorf("不支持的map key，目前 map key 只支持 int32 和 string. %T", mapK)
				}
			default:
				return fmt.Errorf("不支持的map key，目前 map key 只支持 zint32 和 string. %T", mapK)
			}
			// getter

		case *types.Named:

		// typName := v.Obj()
		// // Qual automatically imports packages
		// code.Op("*").Qual(
		// 	typName.Pkg().Path(),
		// 	typName.Name(),
		// )
		default:
			return fmt.Errorf("struct field type not handled: %T", v)
		}
	}
	return nil
}
