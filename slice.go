package main

import (
	"fmt"
	"go/types"

	. "github.com/dave/jennifer/jen"
)

func writeSlice(f *File, v *types.Slice) (string, error) {
	attrTypName := "Slice"

	// 1. 对 struct 做一些准备工作

	// 生成的Map 名字 KV{Key}{Val}
	structName := SliceTypeName(v)

	// type 名字
	valTypStr := getTypString(v.Elem())
	valTyp := v.Elem()
	_, isBasicVal := valTyp.(*types.Basic)

	// 一些预设的类型或者关键字
	// *attr.StrMap
	attrStrMap := func() *Statement { return Id("*").Qual("gitlab.gamesword.com/nut/entitygen/attr", attrTypName) }
	// attr.Field
	// attrField := func() *Statement { return Qual("gitlab.gamesword.com/nut/entitygen/attr", "Field") }
	// 将 name 变量转为 *attr.StrMap类型: (*attr.StrMap)(name)
	convertAttrStrMap := func(name string) *Statement { return Parens(attrStrMap()).Parens(Id(name)) }
	// a *XXXDef
	thisFn := func() *Statement { return Id(thisKeyword).Op("*").Id(structName) }
	// 将 "a" 转为 *attr.StrMap 类型：(*attr.StrMap)(a)
	convertThisFn := func() *Statement { return convertAttrStrMap(thisKeyword) }

	// 3. 写定义  type XXXDef attr.StrMap
	f.Type().Id(structName).Qual("gitlab.gamesword.com/nut/entitygen/attr", attrTypName)

	// 4. 写构造函数
	// EmptyXXXX 和 NewXXX

	// 写 EmptyXXX
	f.Func().Id(EmptyCtor(structName)).Params().Op("*").Id(structName).
		Block(
			Return(Id(NormalCtor(structName)).CallFunc(func(g *Group) {
				g.Nil()
			})),
		)
	// 写 NewXXX
	f.Func().Id(NormalCtor(structName)).ParamsFunc(func(g *Group) {
		g.Id("items").Index().Id(valTypStr)
	}).Op("*").Id(structName).
		BlockFunc(func(g *Group) {
			g.Var().Id("convertData").Index().Interface().Op("=").Index().Interface().Block()
			g.For().Id("k").Op(",").Id("v").Op(":=").Range().Id("items").BlockFunc(
				func(ig *Group) {
					ig.Id("convertData").Index(Id("k")).Op("=").Id("v")
				},
			)

			g.Return(Parens(Op("*").Id(structName)).Params(Qual("gitlab.gamesword.com/nut/entitygen/attr", fmt.Sprintf("New%s", attrTypName)).Call(Id("convertData"))))
		})

	// 5. 写所有字段的 getter/setter
	writeSliceGetSetDel(f, valTypStr, valTyp, isBasicVal, thisFn, convertThisFn)

	// // 6. 写自定义方法
	// // 写 setParent ForEach Equal
	// writeMapCustomMethod(f, structName, keyTypStr, valTypStr, attrField, thisFn, convertThisFn, convertAttrStrMap)

	// // 7. 写 marshal & unmarshal
	// writeMapEncodeDecode(f, keyTypStr, valTypStr, isBasicVal, thisFn, convertThisFn)
	return structName, nil
}

func writeSliceGetSetDel(
	f *File,
	valTypStr string,
	valTyp types.Type,
	isBasicVal bool,
	thisFn func() *Statement,
	convertThisFn func() *Statement,
) {
	// 写 Set
	f.Func().Params(thisFn()).Id("Set").Params(Id("idx").Int(), Id("item").Add(Id(valTypStr))).
		BlockFunc(func(g *Group) {
			if !isBasicVal {
				g.Add(setSliceParentCode("idx", "item", convertThisFn))
			}
			g.Add(convertThisFn().Dot("Set").Call(Id("idx"), Id("item")))
		})

	// 写 Add
	f.Func().Params(thisFn()).Id("Add").Params(Id("item").Add(Id(valTypStr))).
		BlockFunc(func(g *Group) {
			if !isBasicVal {
				g.Id("idx").Op(":=").Id(thisKeyword).Dot("Count").Call()
				g.Add(setSliceParentCode("idx", "item", convertThisFn))
			}
			g.Add(convertThisFn().Dot("Add").Call(Id("item")))
		})

	// 写 At
	attrGetter, shouldReturnConvert := getFieldAttrGetterFnName(valTyp)
	f.Func().Params(thisFn()).Id("At").Params(Id("idx").Int()).Parens(Id(valTypStr)).
		BlockFunc(func(g *Group) {
			g.Id("val").Op(":=").Add(convertThisFn()).Dot(attrGetter).Call(Id("idx"))
			// g.If(Id("val").Op("==").Nil()).Block(
			// 	Return(getNilValue(valTyp), False()),
			// )
			if shouldReturnConvert {
				g.Return(Id("val").Dot("").Parens(Id(valTypStr))) // 做类型转换
			} else {
				g.Return(Id("val"))
			}
		})

	// 写 Delete
	f.Func().Params(thisFn()).Id("DelAt").Params(Id("idx").Int()).Bool().
		Block(
			Return(convertThisFn().Dot("DeleteAt").Call(Id("idx"))),
		)

		// 写 Count
	f.Func().Params(thisFn()).Id("Count").Params().Int().
		Block(
			Return(convertThisFn().Dot("Len").Call()),
		)
}
