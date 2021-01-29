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

func writeMap(f *File, v *types.Map) error {
	err := checkMapKey(v)
	if err != nil {
		return err
	}

	attrTypName := ""
	basicInfo := v.Key().(*types.Basic)
	if basicInfo.Kind() == types.Int32 {
		attrTypName = "Int32Map"
	} else if basicInfo.Kind() == types.String {
		attrTypName = "StrMap"
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
	// *attr.StrMap
	attrStrMap := func() *Statement { return Id("*").Qual("gitlab.gamesword.com/nut/entitygen/attr", attrTypName) }
	// attr.Field
	attrField := func() *Statement { return Qual("gitlab.gamesword.com/nut/entitygen/attr", "Field") }
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
		g.Id("data").Map(Id(keyTypStr)).Id(valTypStr)
	}).Op("*").Id(structName).
		BlockFunc(func(g *Group) {
			g.Var().Id("convertData").Map(Id(keyTypStr)).Interface().Op("=").Map(Id(keyTypStr)).Interface().Block()
			g.For().Id("k").Op(",").Id("v").Op(":=").Range().Id("data").BlockFunc(
				func(ig *Group) {
					ig.Id("convertData").Index(Id("k")).Op("=").Id("v")
				},
			)

			g.Return(Parens(Op("*").Id(structName)).Params(Qual("gitlab.gamesword.com/nut/entitygen/attr", fmt.Sprintf("New%s", attrTypName)).Call(Id("convertData"))))
		})

	// 5. 写所有字段的 getter/setter
	writeMapGetSetDel(f, keyTypStr, valTypStr, valTyp, isBasicVal, thisFn, convertThisFn)

	// 6. 写自定义方法
	// 写 setParent ForEach Equal
	writeParentForEachEqual(f, structName, keyTypStr, valTypStr, attrField, thisFn, convertThisFn, convertAttrStrMap)

	// 7. 写 marshal & unmarshal
	writeMapEncodeDecode(f, keyTypStr, valTypStr, isBasicVal, thisFn, convertThisFn)
	return nil
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
		parentKey = Qual("fmt", "Sprintf").Call(Lit("idx%d").Op(",").Id(keyParamName))
	}
	return Id(valParamName).Dot("setParent").Call(parentKey, convertThisFn())
}
