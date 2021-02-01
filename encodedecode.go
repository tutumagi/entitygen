package main

import (
	. "github.com/dave/jennifer/jen"
)

func writeEncodeDecode(f *File, thisFn func() *Statement, convertThisFn func() *Statement, attrMetaName string) {
	writer := func(encodeFnName string, encodePackageFn *Statement, decodeFnName string, decodePackageFn *Statement) {
		// marshal
		f.Func().Params(thisFn()).Id(encodeFnName).Params().Params(Index().Byte(), Error()).
			Block(
				Return(encodePackageFn.Call(convertThisFn().Dot("ToMap").Call())),
			)
		// unmarshal
		f.Func().Params(thisFn()).Id(decodeFnName).Params(Id("b").Index().Byte()).Error().
			BlockFunc(func(g *Group) {
				g.Id("mm").Id(",").Id("err").Op(":=").Add(decodePackageFn.Call(Id("b")))
				g.If(Id("err").Op("!=").Nil()).Block(
					Return(Id("err")),
				)
				g.Add(convertThisFn().Dot("SetData").Params(Id("mm")))
				g.Add(convertThisFn().Dot("ForEach").Params(
					Func().Params(Id("k").String(), Id("v").Interface()).Bool().
						BlockFunc(func(g *Group) {
							g.If(Id("k").Op("!=").Lit("id").Op("&&").Op("!").Id(attrMetaName).Dot("GetDef").Params(Id("k")).Dot("IsPrimary").Params().Block(
								Id("v").Dot("").Parens(Id("IField")).Dot("setParent").Params(Id("k"), convertThisFn()),
							))
							g.Return(True())
						}),
				))
				g.Return(Nil())
			},
			)
	}
	writer(
		"MarshalJSON",
		Qual("encoding/json", "Marshal"),
		"UnmarshalJSON",
		Id(attrMetaName).Dot("UnmarshalJson"),
	)

	writer(
		"MarshalBSON",
		Qual("go.mongodb.org/mongo-driver/bson", "Marshal"),
		"UnmarshalBSON",
		Id(attrMetaName).Dot("UnmarshalBson"),
	)
}

func writeMapEncodeDecode(
	f *File,
	keyTypStr string,
	valTypStr string,
	isBasicVal bool,
	thisFn func() *Statement,
	convertThisFn func() *Statement,
) {

	writer := func(encodeFnName string, encodePackageFn *Statement, decodeFnName string, decodePackageFn *Statement) {
		// marshal
		f.Func().Params(thisFn()).Id(encodeFnName).Params().Params(Index().Byte(), Error()).
			Block(
				Return(encodePackageFn.Call(convertThisFn().Dot("ToMap").Call())),
			)
		// unmarshal
		f.Func().Params(thisFn()).Id(decodeFnName).Params(Id("b").Index().Byte()).Error().
			BlockFunc(func(g *Group) {
				g.Id("dd").Op(":=").Map(Id(keyTypStr)).Id(valTypStr).Block()
				g.Id("err").Op(":=").Add(decodePackageFn).Call(Id("b"), Op("&").Id("dd"))
				g.If(Id("err").Op("!=").Nil()).Block(
					Return(Id("err")),
				)

				g.Id("convertData").Op(":=").Map(Id(keyTypStr)).Interface().Block()

				g.For().Id("k").Op(",").Id("v").Op(":=").Range().Id("dd").BlockFunc(
					func(ig *Group) {
						// val 不是基础类型，就需要设置一下 parent
						if !isBasicVal {
							ig.Add(setParenctCode("k", "v", keyTypStr, convertThisFn))
						}
						ig.Id("convertData").Index(Id("k")).Op("=").Id("v")
					},
				)
				g.Add(convertThisFn().Dot("SetData").Params(Id("convertData")))

				g.Return(Nil())
			},
			)
	}
	writer(
		"MarshalJSON",
		Qual("encoding/json", "Marshal"),
		"UnmarshalJSON",
		Qual("encoding/json", "Unmarshal"),
	)

	writer(
		"MarshalBSON",
		Qual("go.mongodb.org/mongo-driver/bson", "Marshal"),
		"UnmarshalBSON",
		Qual("go.mongodb.org/mongo-driver/bson", "Unmarshal"),
	)
}

func writeSliceEncodeDecode(
	f *File,
	valTypStr string,
	isBasicVal bool,
	thisFn func() *Statement,
	convertThisFn func() *Statement,
) {

	writer := func(encodeFnName string, encodePackageFn *Statement, decodeFnName string, decodePackageFn *Statement) {
		// marshal
		f.Func().Params(thisFn()).Id(encodeFnName).Params().Params(Index().Byte(), Error()).
			Block(
				Return(encodePackageFn.Call(Id(thisKeyword).Dot("data").Call())),
			)
		// unmarshal
		f.Func().Params(thisFn()).Id(decodeFnName).Params(Id("b").Index().Byte()).Error().
			BlockFunc(func(g *Group) {
				g.Id("dd").Op(":=").Index().Id(valTypStr).Block()
				g.Id("err").Op(":=").Add(decodePackageFn).Call(Id("b"), Op("&").Id("dd"))
				g.If(Id("err").Op("!=").Nil()).Block(
					Return(Id("err")),
				)

				g.Id("convertData").Op(":=").Index().Interface().Block()

				g.For().Id("k").Op(",").Id("v").Op(":=").Range().Id("dd").BlockFunc(
					func(ig *Group) {
						// val 不是基础类型，就需要设置一下 parent
						if !isBasicVal {
							ig.Add(setSliceParentCode("k", "v", convertThisFn))
						}
						ig.Id("convertData").Index(Id("k")).Op("=").Id("v")
					},
				)
				g.Add(convertThisFn().Dot("SetData").Params(Id("convertData")))

				g.Return(Nil())
			},
			)
	}
	writer(
		"MarshalJSON",
		Qual("encoding/json", "Marshal"),
		"UnmarshalJSON",
		Qual("encoding/json", "Unmarshal"),
	)

	writer(
		"MarshalBSON",
		Qual("go.mongodb.org/mongo-driver/bson", "Marshal"),
		"UnmarshalBSON",
		Qual("go.mongodb.org/mongo-driver/bson", "Unmarshal"),
	)
}
