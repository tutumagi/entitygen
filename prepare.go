package main

import (
	"fmt"
	"go/types"
	"reflect"
	"strings"

	. "github.com/dave/jennifer/jen"
)

// 通过原始 struct 名字，获取生成的 struct 名字
func genStructName(srcName string) string {
	return srcName + "Def"
}

// 通过原始 struct 名字，获取生成的 struct 的 meta 变量名字
func genMetaName(srcName string) string {
	return strings.ToLower(srcName) + "AttrDef"
}

type structField struct {
	name       string
	key        string
	typ        types.Type
	storeDB    bool
	base       bool
	cell       bool
	client     bool
	getter     Code
	setter     Code
	setParam   Code
	attrGetter string
	// 该字段类型 转换后的字符串
	// 比如 string 就是 string， int8 就是  int8
	// 自定义类型就要转一下 加一个 Def 后面，比如 Desk 就是 DeskDef
	typString string
	// zero value 对应的 Code
	emptyValue Code
}

func getStructFields(structType *types.Struct) []*structField {
	result := make([]*structField, 0, structType.NumFields())
	for i := 0; i < structType.NumFields(); i++ {
		field := structType.Field(i)
		name := field.Name()

		if name == "id" {
			// NOTE: 目前生成的代码里面跳过 id 的处理
			continue
		}

		typ := field.Type()
		storeDB := false
		flagBase := true // 目前的实现里面属性肯定会存储在 base 里面
		flagCell := false
		client := false

		tagValue := reflect.StructTag(structType.Tag(i))
		key, ok := tagValue.Lookup("key")
		if !ok {
			failErr(fmt.Errorf("field:%s 必须有tag:key", name))
		}
		{
			storeDBStr, ok := tagValue.Lookup("storedb")
			if !ok {
				failErr(fmt.Errorf("field:%s 必须有tag:storedb", name))
			}
			if storeDBStr != "true" && storeDBStr != "false" {
				failErr(fmt.Errorf("field:%s storedb(%s) 必须是 true 或者 false", name, storeDBStr))
			}
			if storeDBStr == "true" {
				storeDB = true
			}
		}

		{
			clientStr, ok := tagValue.Lookup("client")
			if !ok {
				failErr(fmt.Errorf("field:%s 必须有tag:client", name))
			}
			if clientStr != "true" && clientStr != "false" {
				failErr(fmt.Errorf("field:%s client(%s) 必须是 true 或者 false", name, clientStr))
			}
			if clientStr == "true" {
				client = true
			}
		}

		{
			flagStr, ok := tagValue.Lookup("flag")
			if !ok {
				failErr(fmt.Errorf("field:%s 必须有tag:flag", name))
			}
			if flagStr != "base" && flagStr != "cell" {
				failErr(fmt.Errorf("field:%s flag(%s) 必须是 base 或者 cell", name, flagStr))
			}
			if flagStr == "cell" {
				flagCell = true
			}
		}

		result = append(result, &structField{
			name:       name,
			key:        key,
			typ:        typ,
			storeDB:    storeDB,
			base:       flagBase,
			cell:       flagCell,
			client:     client,
			emptyValue: getEmptyValue(typ),
			getter:     Id(fmt.Sprintf("Get%s", name)),
			setter:     Id(fmt.Sprintf("Set%s", name)),
			setParam:   Id(key).Id(getTypString(typ)),
			typString:  getTypString(typ),
			attrGetter: getFieldAttrGetterFnName(typ),
		})
	}
	return result
}

func getEmptyValue(typ types.Type) Code {
	switch v := typ.(type) {
	case *types.Basic:
		switch v.Kind() {
		case types.String, types.UntypedString:
			return Lit("")
		case types.Int, types.Uint, types.Int8, types.Uint8, types.Int16, types.Uint16, types.Int32, types.Uint32, types.Int64, types.Uint64, types.Float32, types.Float64:
			return Lit(0)
		case types.Bool:
			return Lit(false)
		default:
			return Nil()
		}
	default:
		return Nil()
	}
}

func getTypString(typ types.Type) string {
	switch v := typ.(type) {
	case *types.Basic:
		return v.String()

	case *types.Pointer:
		// types.Pointer 就用 .Elem 解引用
		// types.Named 就用 .Underlying 获取引用的类型
		switch vv := v.Elem().(type) {
		case *types.Basic:
			return fmt.Sprintf("*%s", vv.String())
		case *types.Struct:
			return fmt.Sprintf("*%s", vv.String())
		case *types.Named:
			switch vvv := vv.Underlying().(type) {
			case *types.Basic:
				return fmt.Sprintf("*%s", vvv.String())
			case *types.Struct:
				// 这样就不会带包名和路径名，否则会出现 entitygen/entitydef.Desk 这种情况
				return fmt.Sprintf("*%s", genStructName(vv.Obj().Name()))
				// return fmt.Sprintf("*%s", vvv.String())
			default:
				failErr(fmt.Errorf("1 不支持的类型 %s", vvv.String()))
			}
		default:
			failErr(fmt.Errorf("2 不支持的类型 %s", vv.String()))
		}
	default:
		failErr(fmt.Errorf("3 不支持的类型 %s", v.String()))
	}
	return ""
}

// 获取 attr.StrMap 或者 attr.Int32Map 的 getter 方法名
func getFieldAttrGetterFnName(typ types.Type) string {
	switch v := typ.(type) {
	case *types.Basic:
		// attr.StrMap 的 get 方法
		// 如果是基础类型，则直接大写第一个字母的方法进行 getter 比如 int32 就是 .Int32("xxx")
		// 如果是 string 类型，则使用 Str 方法，比如 .Str("yyy")
		attrGetFuncName := strings.Title(v.Name())
		switch v.Kind() {
		case types.String, types.UntypedString:
			attrGetFuncName = "Str"
		}
		return attrGetFuncName
	default:
		return "Value"
	}
}
