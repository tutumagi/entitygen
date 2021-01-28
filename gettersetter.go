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
			f.Func().Params(thisFn()).Add(field.getter).Params().Id(field.typName).
				Block(
					ReturnFunc(func(g *Group) {
						statement := g.Add(convertThisFn()).Dot(field.attrGetter).Params(Lit(field.key))
						// 如果不是基础类型，则加上类型转换
						if !isBasic {
							statement.Dot("").Parens(Id(field.typName))
						}
					}),
				)

			//  写 setter
			f.Func().Params(thisFn()).Add(field.setter).Params(field.setParam).
				BlockFunc(func(g *Group) {
					if !isBasic {
						g.Id(field.key).Dot("setParent").Call(Lit(field.key), convertThisFn())
					}
					g.Add(convertThisFn()).Dot("Set").Params(Lit(field.key), Id(field.key))
				})

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

func writeMapGetSetDel(
	f *File,
	keyTypStr string,
	valTypStr string,
	valTyp types.Type,
	isBasicVal bool,
	thisFn func() *Statement,
	convertThisFn func() *Statement,
) {
	// 写 Set
	f.Func().Params(thisFn()).Id("Set").Params(Id("k").Add(Id(keyTypStr)), Id("v").Add(Id(valTypStr))).
		BlockFunc(func(g *Group) {
			if !isBasicVal {
				g.Add(setParenctCode("k", "v", keyTypStr, convertThisFn))
			}
			g.Add(convertThisFn().Dot("Set").Call(Id("k"), Id("v")))
		})

	// 写 Get
	attrGetter, shouldReturnConvert := getFieldAttrGetterFnName(valTyp)
	f.Func().Params(thisFn()).Id("Get").Params(Id("k").Add(Id(keyTypStr))).Id(valTypStr).
		BlockFunc(func(g *Group) {
			statement := g.Return(Add(convertThisFn()).Dot(attrGetter).Call(Id("k")))
			if shouldReturnConvert {
				statement.Dot("").Parens(Id(valTypStr)) // 做类型转换
			}
		})

	// 写 Delete
	f.Func().Params(thisFn()).Id("Delete").Params(Id("k").Add(Id(keyTypStr))).Bool().
		Block(
			Return(convertThisFn().Dot("Delete").Call(Id("k"))),
		)
}
