package main

import (
	"fmt"
	"go/types"

	. "github.com/dave/jennifer/jen"
)

func writeGetterSetter(f *File, entInfo *entStructInfo, fields []*fieldInfo, thisFn func() *Statement, convertThisFn func() *Statement) error {
	// 写自定义属性的 getter/setter
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
					// if !isBasic {
					// 	const tmp = "tmp"
					// 	// g.If(Id(field.key).Op("==").Nil()).Block(Return())
					// 	// // 先判断下当前对象有没有这个字段
					// 	// g.Id(tmp).Op(":=").Id(thisKeyword).Op(".").Add(field.getter).Params()
					// 	// g.If(Id(tmp).Op("==").Nil()).BlockFunc(func(ggg *Group) {
					// 	// 	// 如果没有，就用参数拷贝构造一个新的值，赋值给当前对象
					// 	// 	ggg.Id(tmp).Op("=").Op(CopyCtor(trimHeadStar(field.typName))).Call(Id(field.key))
					// 	// }).Else().BlockFunc(func(ggg *Group) {
					// 	// 	// 如果当前对象有这个字段，则直接更新即可
					// 	// 	ggg.Id(tmp).Dot(updateFuncName).Call(Id(field.key))
					// 	// })
					// 	g.Id(tmp).Op(":=").Op(CopyCtor(trimHeadStar(field.typName))).Call(Id(field.key))

					// 	g.Id(tmp).Dot(setParentFuncName).Call(Lit(field.key), convertThisFn())
					// 	g.Add(convertThisFn()).Dot("Set").Params(Lit(field.key), Id(tmp))
					// } else {
					// 	g.Add(convertThisFn()).Dot("Set").Params(Lit(field.key), Id(field.key))
					// }
					if !isBasic {
						g.Id(field.key).Dot(setParentFuncName).Call(Lit(field.key), convertThisFn())
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

	// 写实体内置的 getter/setter

	writeBuiltinProp(f, entInfo, thisFn, convertThisFn)

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
				// g.Id("v").Op("=").Op(CopyCtor(trimHeadStar(valTypStr))).Call(Id("v"))

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

func writeBuiltinProp(
	f *File,
	entInfo *entStructInfo,
	thisFn func() *Statement,
	convertThisFn func() *Statement,
) {
	if entInfo == nil {
		return
	}
	if entInfo.rot != nil {
		// GetRot
		f.Func().Params(thisFn()).Id("GetRot").Params().Op("*").Add(attrVec3()).
			BlockFunc(func(g *Group) {
				g.Id("val").Op(":=").Add(convertThisFn()).Dot("Value").Call(Lit(buildinRotKey))

				g.If(Id("val").Op("==").Nil()).Block(Return(Nil()))
				g.Return(Id("val").Dot("").Parens(Op("*").Add(attrVec3()))) // 做类型转换
			})

		//  SetRot
		f.Func().Params(thisFn()).Id("SetRot").Params(Id(buildinRotKey).Op("*").Add(attrVec3())).
			BlockFunc(func(g *Group) {
				g.Id(buildinRotKey).Dot(setParentFuncName).Call(Lit(buildinRotKey), convertThisFn())
				g.Add(convertThisFn()).Dot("Set").Params(Lit(buildinRotKey), Id(buildinRotKey))
			})
	}

	if entInfo.pos != nil {
		// GetPos
		f.Func().Params(thisFn()).Id("GetPos").Params().Op("*").Add(attrVec3()).
			BlockFunc(func(g *Group) {
				g.Id("val").Op(":=").Add(convertThisFn()).Dot("Value").Call(Lit(buildinPosKey))

				g.If(Id("val").Op("==").Nil()).Block(Return(Nil()))
				g.Return(Id("val").Dot("").Parens(Op("*").Add(attrVec3()))) // 做类型转换
			})

		//  SetPos
		f.Func().Params(thisFn()).Id("SetPos").Params(Id(buildinPosKey).Op("*").Add(attrVec3())).
			BlockFunc(func(g *Group) {
				g.Id(buildinPosKey).Dot(setParentFuncName).Call(Lit(buildinPosKey), convertThisFn())
				g.Add(convertThisFn()).Dot("Set").Params(Lit(buildinPosKey), Id(buildinPosKey))
			})
	}

	// GetId
	f.Func().Params(thisFn()).Id("GetId").Params().String().
		BlockFunc(func(g *Group) {
			g.Return(convertThisFn().Dot("Str").Call(Lit(buildinIDKey)))
		})

	//  SetId
	f.Func().Params(thisFn()).Id("SetId").Params(Id(buildinIDKey).String()).
		BlockFunc(func(g *Group) {
			g.Add(convertThisFn()).Dot("Set").Params(Lit(buildinIDKey), Id(buildinIDKey))
		})
}
