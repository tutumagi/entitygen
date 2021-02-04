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
		case *types.Basic, *types.Pointer, *types.Map, *types.Slice:
			// 写 getter
			_, isBasic := v.(*types.Basic)
			if vvm, ok := v.(*types.Map); ok {
				if err := checkMapKey(vvm); err != nil {
					return err
				}
			}
			// func (a *XXXDef) GetField() FieldType
			f.Func().Params(thisFn()).Add(field.getter).Params().Id(field.typName).
				BlockFunc(func(g *Group) {
					g.Id("val").Op(":=").Add(convertThisFn()).Dot(field.attrGetter).Call(Lit(field.key))
					if !isBasic {
						g.If(Id("val").Op("==").Nil()).Block(Return(Nil()))
						g.Return(Id("val").Dot("").Parens(Id(field.typName))) // 做类型转换
					} else {
						g.Return(Id("val"))
					}
				})

			//  写 setter
			f.Func().Params(thisFn()).Add(field.setter).Params(field.setParam).
				BlockFunc(func(g *Group) {
					if !isBasic {
						// g.If(Id(field.key).Op("==").Nil()).Block(Return())
						g.Id(field.key).Dot("setParent").Call(Lit(field.key), convertThisFn())
						g.Add(convertThisFn()).Dot("Set").Params(Lit(field.key), Id(field.key))
					} else {
						g.Add(convertThisFn()).Dot("Set").Params(Lit(field.key), Id(field.key))
					}

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
				// g.If(Id("v").Op("==").Nil()).Block(Return())
				g.Add(setParenctCode("k", "v", keyTypStr, convertThisFn))
				g.Add(convertThisFn().Dot("Set").Call(Id("k"), Id("v")))
			} else {
				g.Add(convertThisFn().Dot("Set").Call(Id("k"), Id("v")))
			}
		})

	// 写 Get
	attrGetter, shouldReturnConvert := getFieldAttrGetterFnName(valTyp)
	f.Func().Params(thisFn()).Id("Get").Params(Id("k").Add(Id(keyTypStr))).Id(valTypStr).
		BlockFunc(func(g *Group) {
			g.Id("val").Op(":=").Add(convertThisFn()).Dot(attrGetter).Call(Id("k"))
			if shouldReturnConvert {
				g.If(Id("val").Op("==").Nil()).Block(Return(Nil()))
				g.Return(Id("val").Dot("").Parens(Id(valTypStr))) // 做类型转换
			} else {
				g.Return(Id("val"))
			}
		})

	// 写 Delete
	f.Func().Params(thisFn()).Id("Delete").Params(Id("k").Add(Id(keyTypStr))).Bool().
		Block(
			Return(convertThisFn().Dot("Delete").Call(Id("k"))),
		)

		// 写 Count
	f.Func().Params(thisFn()).Id("Count").Params().Int().
		Block(
			Return(convertThisFn().Dot("Len").Call()),
		)
}
