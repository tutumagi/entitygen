package main

import (
	"fmt"
	"go/types"

	. "github.com/dave/jennifer/jen"
)

func checkMapKey(v *types.Map) error {
	switch mapK := v.Key().(type) {
	case *types.Basic:
		if mapK.Kind() == types.Int32 || mapK.Kind() == types.String {
		} else {
			return fmt.Errorf("不支持的map key，目前 map key 只支持 int32 和 string. %T", mapK)
		}
	default:
		return fmt.Errorf("不支持的map key，目前 map key 只支持 zint32 和 string. %T", mapK)
	}
	return nil
}

func writeMap(f *File, v *types.Map) (string, error) {
	err := checkMapKey(v)
	if err != nil {
		return "", err
	}

	// 1. 对 struct 做一些准备工作

	// 生成的Map 名字 KV{Key}{Val}
	structName := MapTypeName(v)

	// key type 名字
	keyTypStr := v.Key().String()
	valTypStr := getTypString(v.Elem())
	valTyp := v.Elem()
	_, isBasicVal := valTyp.(*types.Basic)

	// 一些预设的类型或者关键字

	attrTypName := ""
	basicInfo := v.Key().(*types.Basic)
	if basicInfo.Kind() == types.Int32 {
		attrTypName = "Int32Map"
	} else if basicInfo.Kind() == types.String {
		attrTypName = "StrMap"
	} else {
		failErr(fmt.Errorf("不支持 int32,string 之外作为 key 的 map(key:%s)", basicInfo.Name()))
	}

	attrType, thisFn, convertThisFn, convertAttrType := aboutThisCode(structName, attrTypName)

	// 3. 写定义  type XXXDef attr.StrMap
	f.Type().Id(structName).Add(attrType())

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
		g.Id("data").Map(Id(keyTypStr)).Id(valTypStr)
	}).Op("*").Id(structName).
		BlockFunc(func(g *Group) {
			g.Var().Id("convertData").Map(Id(keyTypStr)).Interface().Op("=").Map(Id(keyTypStr)).Interface().Block()
			g.For().Id("k").Op(",").Id("v").Op(":=").Range().Id("data").BlockFunc(
				func(ig *Group) {
					ig.Id("convertData").Index(Id("k")).Op("=").Id("v")
				},
			)

			g.Return(Parens(Op("*").Id(structName)).Params(Qual(attrPackageName, fmt.Sprintf("New%s", attrTypName)).Call(Id("convertData"))))
		})
	// 拷贝构造
	f.Func().Id(CopyCtor(structName)).Params(Id("value").Op("*").Id(structName)).Op("*").Id(structName).
		BlockFunc(func(g *Group) {
			g.If(Id("value").Op("==").Nil()).Block(Return(Nil()))
			g.Id("a").Op(":=").Id(EmptyCtor(structName)).Call()
			g.Id("value").Dot("ForEach").Call(Func().Params(Id("k").Id(keyTypStr), Id("v").Id(valTypStr)).Bool().BlockFunc(func(ggg *Group) {
				if isBasicVal {
					ggg.Id("a").Dot("Set").Call(Id("k"), Id("v"))
				} else {
					ggg.Id("a").Dot("Set").Call(Id("k"), Id(CopyCtor(trimHeadStar(valTypStr))).Call(Id("v")))
				}
				ggg.Return(True())
			}))

			g.Return(Id("a"))
		})

		// update
	// f.Func().Params(thisFn()).Id(updateFuncName).Params(Id("value").Op("*").Id(structName)).
	// 	BlockFunc(func(g *Group) {
	// 		g.If(Id("value").Op("==").Nil()).Block(Return())
	// 		g.Add(convertThisFn()).Dot("Clear").Call()
	// 		g.Id("value").Dot("ForEach").Call(Func().Params(Id("k").Id(keyTypStr), Id("v").Id(valTypStr)).Bool().BlockFunc(func(ggg *Group) {
	// 			ggg.Id(thisKeyword).Dot("Set").Call(Id("k"), Id("v"))
	// 			ggg.Return(True())
	// 		}))
	// 	})

	// 5. 写所有字段的 getter/setter
	writeMapGetSetDel(f, keyTypStr, valTypStr, valTyp, isBasicVal, thisFn, convertThisFn)

	// 6. 写自定义方法
	// 写 setParent ForEach Equal
	writeMapCustomMethod(f, structName, attrType, keyTypStr, valTypStr, thisFn, convertThisFn, convertAttrType)

	// 7. 写 marshal & unmarshal
	writeMapEncodeDecode(f, keyTypStr, valTypStr, isBasicVal, thisFn, convertThisFn)
	return structName, nil
}

func setParenctCode(
	keyParamName string,
	valParamName string,
	keyTypStr string,
	convertThisFn func() *Statement,
) *Statement {
	parentKey := Id(keyParamName)
	// NOTE: 这里写死了 key 是 int32 的类型
	if keyTypStr == "int32" {
		// mk 表示 map key :)
		parentKey = Qual("fmt", "Sprintf").Call(Lit("mk%d").Op(",").Id(keyParamName))
	}
	return Id(valParamName).Dot(setParentFuncName).Call(parentKey, convertThisFn())
}

func setSliceParentCode(
	idxParamName string,
	valParamName string,
	convertThisFn func() *Statement,
) *Statement {
	parentKey := Qual("fmt", "Sprintf").Call(Lit("ik%d").Op(",").Id(idxParamName))

	return Id(valParamName).Dot(setParentFuncName).Call(parentKey, convertThisFn())
}
