package main

import (
	"go/types"

	. "github.com/dave/jennifer/jen"
)

func writeStruct(f *File, sourceTypeName string, structType *types.Struct) {
	// 1. 对 struct 做一些准备工作
	// 读取 types.Struct 所有字段信息，计算出我们要的信息，并做合法性判断
	fields := getStructFields(structType)

	// 生成的结构体名字 XXXDef
	structName := genStructName(sourceTypeName)
	// 生成的对应的数据结构描述的名字 XXXAttrDef
	attrMetaName := genMetaName(sourceTypeName)

	// 一些预设的类型或者关键字
	// *attr.StrMap
	attrStrMap := func() *Statement { return Id("*").Qual("entitygen/attr", "StrMap") }
	// attr.Field
	attrField := func() *Statement { return Qual("entitygen/attr", "Field") }
	// 将 name 变量转为 *attr.StrMap类型: (*attr.StrMap)(name)
	convertAttrStrMap := func(name string) *Statement { return Parens(attrStrMap()).Parens(Id(name)) }
	// a *XXXDef
	thisFn := func() *Statement { return Id(thisKeyword).Op("*").Id(structName) }
	// 将 "a" 转为 *attr.StrMap 类型：(*attr.StrMap)(a)
	convertThisFn := func() *Statement { return convertAttrStrMap(thisKeyword) }

	// 2. 写 attrDef
	writeAttrMeta(f, attrMetaName, fields)

	// 3. 写定义  type XXXDef attr.StrMap
	f.Type().Id(structName).Qual(
		"entitygen/attr",
		"StrMap",
	)

	// 4. 写构造函数
	writeCtor(f, structName, fields)

	// 5. 写所有字段的 getter/setter
	err := writeGetterSetter(f, fields, thisFn, convertThisFn)
	if err != nil {
		failErr(err)
	}

	// 6. 写自定义方法
	writeCustomMethod(f, structName, attrField, thisFn, convertThisFn, convertAttrStrMap)

	// 7. 写 marshal & unmarshal
	writeEncodeDecode(f, thisFn, convertThisFn, attrMetaName)
}
