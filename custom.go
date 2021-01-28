package main

import (
	. "github.com/dave/jennifer/jen"
)

func writeCustomMethod(
	f *File,
	structName string,
	attrField func() *Statement,
	thisFn func() *Statement,
	convertThisFn func() *Statement,
	convertAttrStrMap func(string) *Statement,
) {

	// 写 changekey 相关的
	writeChangeKey(f, thisFn, convertThisFn)
	// 写 setParent ForEach Equal
	writeParentForEachEqual(f, structName, "string", "interface{}", attrField, thisFn, convertThisFn, convertAttrStrMap)
}

func writeChangeKey(
	f *File,
	thisFn func() *Statement,
	convertThisFn func() *Statement,
) {
	f.Func().Params(thisFn()).Id("HasChange").Params().Bool().
		Block(
			Return(convertThisFn().Dot("HasChange").Call()),
		)

	f.Func().Params(thisFn()).Id("ChangeKey").Params().Map(String()).Struct().
		Block(
			Return(convertThisFn().Dot("ChangeKey").Call()),
		)

	f.Func().Params(thisFn()).Id("ClearChangeKey").Params().
		Block(
			convertThisFn().Dot("ClearChangeKey").Call(),
		)
}

func writeParentForEachEqual(
	f *File,
	structName string,
	keyTyp string,
	valTyp string,
	attrField func() *Statement,
	thisFn func() *Statement,
	convertThisFn func() *Statement,
	convertAttrStrMap func(string) *Statement,
) {
	// 4. 写 setParent
	f.Func().Params(thisFn()).Id("setParent").Params(Id("k").String(), Id("parent").Add(attrField())).
		Block(
			convertThisFn().Dot("SetParent").Call(Id("k"), Id("parent")),
		)

	// 5. ForEach
	// func(k [KeyType], v [ValType])bool
	forEachParamSign := Func().Params(Id("k").Add(Id(keyTyp)), Id("v").Add(Id(valTyp))).Bool()
	forEachUnderlyingSign := Func().Params(Id("k").Add(Id(keyTyp)), Id("v").Interface()).Bool()
	f.Func().Params(thisFn()).Id("ForEach").Params(Id("fn").Add(forEachParamSign)).
		BlockFunc(func(g *Group) {
			statement := g.Add(convertThisFn()).Dot("ForEach")
			if valTyp == "interface{}" {
				// 如果 val 是 interface{} ，则直接 call 底层 map 的 ForEach 方法
				statement.Call(Id("fn"))
			} else {
				// 否则做一层类型转换
				statement.Call(forEachUnderlyingSign.Block(
					Return(Id("fn").Call(Id("k"), Id("v").Dot("").Parens(Id(valTyp)))),
				))
			}
		})

	// 写 Equal
	f.Func().Params(thisFn()).Id("Equal").Params(Id("other").Op("*").Id(structName)).Bool().Block(
		Return(convertThisFn().Dot("Equal").Call(convertAttrStrMap("other"))),
	)
}
