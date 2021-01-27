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
		case *types.Basic, *types.Pointer, *types.Map:
			// 写 getter
			_, isBasic := v.(*types.Basic)
			if vvm, ok := v.(*types.Map); ok {
				if err := checkMapKey(vvm); err != nil {
					return err
				}
			}
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
