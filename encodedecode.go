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
	// // marshal
	// f.Func().Params(thisFn()).Id("MarshalJSON").Params().Params(Index().Byte(), Error()).
	// 	Block(
	// 		Return(Qual("encoding/json", "Marshal").Call(convertThisFn().Dot("ToMap").Params())),
	// 	)
	// // unmarshal
	// f.Func().Params(thisFn()).Id("UnmarshalJSON").Params(Id("b").Index().Byte()).Error().
	// 	BlockFunc(func(g *Group) {
	// 		g.Id("mm").Id(",").Id("err").Op(":=").Id(attrMetaName).Dot("UnmarshalJson").Params(Id("b"))
	// 		g.If(Id("err").Op("!=").Nil()).Block(
	// 			Return(Id("err")),
	// 		)
	// 		g.Add(convertThisFn().Dot("SetData").Params(Id("mm")))
	// 		g.Add(convertThisFn().Dot("ForEach").Params(
	// 			Func().Params(Id("k").String(), Id("v").Interface()).Bool().
	// 				BlockFunc(func(g *Group) {
	// 					g.If(Id("k").Op("!=").Lit("id").Op("&&").Op("!").Id(attrMetaName).Dot("GetDef").Params(Id("k")).Dot("IsPrimary").Params().Block(
	// 						Id("v").Dot("").Parens(Id("IField")).Dot("setParent").Params(Id("k"), convertThisFn()),
	// 					))
	// 					g.Return(True())
	// 				}),
	// 		))
	// 		g.Return(Nil())
	// 	},
	// 	)
}
