package main

import (
	. "github.com/dave/jennifer/jen"
)

func writeStructCustomMethod(
	f *File,
	structName string,
	attrType func() *Statement,
	thisFn func() *Statement,
	convertThisFn func() *Statement,
	convertAttrStrMap func(string) *Statement,
) {

	// 写 changekey 相关的
	writeChangeKey(f, thisFn, convertThisFn)
	// 写 setParent ForEach Equal
	writeParentForEachEqual(f, structName, attrType, "string", "interface{}", thisFn, convertThisFn, convertAttrStrMap)

	writeMapData(f, "string", "interface{}", thisFn)
}

func writeMapCustomMethod(
	f *File,
	structName string,
	attrType func() *Statement,
	keyTypStr string,
	valTypStr string,
	thisFn func() *Statement,
	convertThisFn func() *Statement,
	convertAttrStrMap func(string) *Statement,
) {
	writeParentForEachEqual(f, structName, attrType, keyTypStr, valTypStr, thisFn, convertThisFn, convertAttrStrMap)

	writeHas(f, keyTypStr, thisFn, convertThisFn)

	writeMapData(f, keyTypStr, valTypStr, thisFn)
}

func writeSliceCustomMethod(
	f *File,
	structName string,
	attrType func() *Statement,
	keyTypStr string,
	valTypStr string,
	thisFn func() *Statement,
	convertThisFn func() *Statement,
	convertAttrStrMap func(string) *Statement,
) {
	writeParentForEachEqual(f, structName, attrType, keyTypStr, valTypStr, thisFn, convertThisFn, convertAttrStrMap)
	writeSliceData(f, valTypStr, thisFn)
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
	attrType func() *Statement,
	keyTypStr string,
	valTypStr string,
	thisFn func() *Statement,
	convertThisFn func() *Statement,
	convertAttrStrMap func(string) *Statement,
) {
	// 4. 写 setParent
	f.Func().Params(thisFn()).Id(setParentFuncName).Params(Id("k").String(), Id("parent").Add(attrField())).
		Block(
			convertThisFn().Dot("SetParent").Call(Id("k"), Id("parent")),
		)

	// 5. ForEach
	// func(k [KeyType], v [ValType])bool
	forEachParamSign := Func().Params(Id("k").Add(Id(keyTypStr)), Id("v").Add(Id(valTypStr))).Bool()
	forEachUnderlyingSign := Func().Params(Id("k").Add(Id(keyTypStr)), Id("v").Interface()).Bool()
	f.Func().Params(thisFn()).Id("ForEach").Params(Id("fn").Add(forEachParamSign)).
		BlockFunc(func(g *Group) {
			statement := g.Add(convertThisFn()).Dot("ForEach")
			if valTypStr == "interface{}" {
				// 如果 val 是 interface{} ，则直接 call 底层 map 的 ForEach 方法
				statement.Call(Id("fn"))
			} else {
				// 否则做一层类型转换
				statement.Call(forEachUnderlyingSign.Block(
					Return(Id("fn").Call(Id("k"), Id("v").Dot("").Parens(Id(valTypStr)))),
				))
			}
		})

	// 写 Equal
	f.Func().Params(thisFn()).Id("Equal").Params(Id("other").Op("*").Id(structName)).Bool().Block(
		Return(convertThisFn().Dot("Equal").Call(convertAttrStrMap("other"))),
	)

	// 写 undertype
	f.Func().Params(thisFn()).Id("Undertype").Params().Interface().Block(
		Return(convertThisFn()),
	)
}

func writeHas(
	f *File,
	keyTypStr string,
	thisFn func() *Statement,
	convertThisFn func() *Statement,
) {
	// 4. 写 setParent
	f.Func().Params(thisFn()).Id("Has").Params(Id("k").Id(keyTypStr)).Bool().
		Block(
			Return(convertThisFn().Dot("Has").Call(Id("k"))),
		)
}

func writeSliceData(
	f *File,
	valTypStr string,
	thisFn func() *Statement,
) {
	f.Func().Params(thisFn()).Id("Data").Params().Index().Id(valTypStr).
		BlockFunc(func(g *Group) {
			g.Id("dd").Op(":=").Index().Id(valTypStr).Block()
			g.Id(thisKeyword).Dot("ForEach").Call(Func().Params(Id("idx").Int(), Id("v").Id(valTypStr)).Bool().Block(
				Id("dd").Op("=").Append(Id("dd"), Id("v")),

				Return(True()),
			))
			g.Return(Id("dd"))
		})
}

func writeMapData(
	f *File,
	keyTypStr string,
	valTypStr string,
	thisFn func() *Statement,
) {
	f.Func().Params(thisFn()).Id("Data").Params().Map(Id(keyTypStr)).Id(valTypStr).
		BlockFunc(func(g *Group) {
			g.Id("dd").Op(":=").Map(Id(keyTypStr)).Id(valTypStr).Block()
			g.Id(thisKeyword).Dot("ForEach").Call(Func().Params(Id("k").Id(keyTypStr), Id("v").Id(valTypStr)).Bool().Block(
				Id("dd").Index(Id("k")).Op("=").Id("v"),
				Return(True()),
			))
			g.Return(Id("dd"))
		})
}
