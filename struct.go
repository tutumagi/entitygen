package main

import (
	"go/types"
	"strings"

	. "github.com/dave/jennifer/jen"
	"gitlab.gamesword.com/nut/entitygen/attr"
)

func writeStruct(f *File, sourceTypeName string, structType *types.Struct) string {
	// 如果有 Def 后缀，则认为该结构体是实体定义
	isEntity := strings.HasSuffix(sourceTypeName, "Def")

	// 1. 对 struct 做一些准备工作
	// 读取 types.Struct 所有字段信息，计算出我们要的信息，并做合法性判断
	// 结构体类型名字如果是以 Def 结尾则表示是实体类型
	fields := getStructFields(structType, isEntity)

	// 生成的结构体名字
	structName := StructTypeName(sourceTypeName)
	// 生成的对应的数据结构描述的名字 XXXMeta
	attrMetaName := StructMetaName(sourceTypeName)
	if !isEntity {
		attrMetaName = attr.LowerFirst(attrMetaName)
	}

	// a *XXXDef
	attrType, thisFn, convertThisFn, convertAttrType := aboutThisCode(structName, "StrMap")

	// 2. 写 attrDef
	writeAttrMeta(f, attrMetaName, structName, isEntity, fields)

	// 3. 写定义  type XXXDef attr.StrMap
	f.Type().Id(structName).Add(attrType())

	// 4. 写构造函数
	writeCtor(f, structName, isEntity, fields)

	// 5. 写所有字段的 getter/setter
	err := writeGetterSetter(f, isEntity, fields, thisFn, convertThisFn)
	if err != nil {
		failErr(err)
	}

	// 6. 写自定义方法
	writeStructCustomMethod(f, structName, attrType, thisFn, convertThisFn, convertAttrType)

	// 7. 写 marshal & unmarshal
	writeEncodeDecode(f, thisFn, convertThisFn, attrMetaName)

	// 写 id, pos, rot 的 getter setter
	return structName
}
