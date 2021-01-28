package main

import (
	"fmt"
	"go/types"
	"strings"

	. "github.com/dave/jennifer/jen"
)

func genMapTypName(v *types.Map) string {
	key := strings.Title(v.Key().String())
	if key == "String" {
		key = "Str"
	}

	val := getTypString(v.Elem())
	if strings.HasPrefix(val, "*") {
		val = strings.TrimLeft(val, "*")
	}
	val = strings.Title(val)
	if val == "String" {
		val = "Str"
	}
	return fmt.Sprintf("KV%s%s", key, val)
}

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
	structName := genMapTypName(v)

	// key type 名字
	keyTypStr := v.Key().String()
	valTypStr := getTypString(v.Elem())
	valTyp := v.Elem()

	// 一些预设的类型或者关键字
	// *attr.StrMap
	attrStrMap := func() *Statement { return Id("*").Qual("entitygen/attr", attrTypName) }
	// attr.Field
	attrField := func() *Statement { return Qual("entitygen/attr", "Field") }
	// 将 name 变量转为 *attr.StrMap类型: (*attr.StrMap)(name)
	convertAttrStrMap := func(name string) *Statement { return Parens(attrStrMap()).Parens(Id(name)) }
	// a *XXXDef
	thisFn := func() *Statement { return Id(thisKeyword).Op("*").Id(structName) }
	// 将 "a" 转为 *attr.StrMap 类型：(*attr.StrMap)(a)
	convertThisFn := func() *Statement { return convertAttrStrMap(thisKeyword) }

	// 3. 写定义  type XXXDef attr.StrMap
	f.Type().Id(structName).Qual("entitygen/attr", attrTypName)

	// 4. 写构造函数
	// EmptyXXXX 和 NewXXX
	emptyCtorName := fmt.Sprintf("Empty%s", structName)
	normalCtorName := fmt.Sprintf("New%s", structName)
	// 写 EmptyXXX
	f.Func().Id(emptyCtorName).Params().Op("*").Id(structName).
		Block(
			Return(Id(normalCtorName).CallFunc(func(g *Group) {
				g.Nil()
			})),
		)
	// 写 NewXXX
	f.Func().Id(normalCtorName).ParamsFunc(func(g *Group) {
		g.Id("data").Map(Id(keyTypStr)).Id(valTypStr)
	}).Op("*").Id(structName).
		BlockFunc(func(g *Group) {
			g.Var().Id("convertData").Map(Id(keyTypStr)).Interface().Op("=").Map(Id(keyTypStr)).Interface().Block()
			g.For().Id("k").Op(",").Id("v").Op(":=").Range().Id("data").BlockFunc(
				func(ig *Group) {
					ig.Id("convertData").Index(Id("k")).Op("=").Id("v")
				},
			)

			g.Return(Parens(Op("*").Id(structName)).Params(Qual("entitygen/attr", fmt.Sprintf("New%s", attrTypName)).Call(Id("convertData"))))
		})

	// 5. 写所有字段的 getter/setter
	writeMapGetSetDel(f, keyTypStr, valTypStr, valTyp, thisFn, convertThisFn)

	// 6. 写自定义方法
	// 写 setParent ForEach Equal
	writeParentForEachEqual(f, structName, keyTypStr, valTypStr, attrField, thisFn, convertThisFn, convertAttrStrMap)

	// 7. 写 marshal & unmarshal
	// writeEncodeDecode(f, thisFn, convertThisFn, attrMetaName)
	return nil
}
