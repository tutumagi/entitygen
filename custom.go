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
	// 3. 写 changekey 相关的
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

	// 4. 写 setParent
	f.Func().Params(thisFn()).Id("setParent").Params(Id("k").String(), Id("parent").Add(attrField())).
		Block(
			convertThisFn().Dot("SetParent").Call(Id("k"), Id("parent")),
		)

	// 5. ForEach
	f.Func().Params(thisFn()).Id("ForEach").Params(Id("fn").Func().Params(Id("s").String(), Id("v").Interface()).Bool()).
		Block(
			convertThisFn().Dot("ForEach").Call(Id("fn")),
		)

	// 写 Equal
	f.Func().Params(thisFn()).Id("Equal").Params(Id("other").Op("*").Id(structName)).Bool().Block(
		Return(convertThisFn().Dot("Equal").Call(convertAttrStrMap("other"))),
	)
}
