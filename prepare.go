package main

import (
	"fmt"
	"go/types"
	"reflect"
	"strings"
	"unicode"

	. "github.com/dave/jennifer/jen"
)

const thisKeyword = "a"

// 通过原始 struct 名字，获取生成的 struct 名字
func genStructName(srcName string) string {
	return srcName + "Def"
}

// 通过原始 struct 名字，获取生成的 struct 的 meta 变量名字
func genMetaName(srcName string) string {
	return strings.ToLower(srcName) + "Meta"
}

type structField struct {
	name       string // 字段名
	key        string // 字段存储在 map 中的 key，目前的规则是字段名首字母小写
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

			key = func() string {
				for _, c := range name {
					return string(unicode.ToLower(c)) + name[1:]
				}
				return name
			}()
			fmt.Printf("field:%s 没有 key，使用 name 作为 key(%s) \n", name, key)
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

		attrGetter, _ := getFieldAttrGetterFnName(typ)
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
			attrGetter: attrGetter,
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
	// case *types.Struct:
	// 	return v.String()
	case *types.Named:
		// 如果是 命名字段类型，比如 struct { A *Desk }
		// 则 v.Obj().Name() 为 "Desk", v.Underlying() 为 types.Struct(Desk)
		return getNamedTypName(v.Obj().Name(), v.Underlying())
	case *types.Pointer:
		// types.Pointer 就用 .Elem 解引用
		// types.Named 就用 .Underlying 获取引用的类型
		return fmt.Sprintf("*%s", getTypString(v.Elem()))
	case *types.Map:
		return fmt.Sprintf("*%s", genMapTypName(v))
	default:
		failErr(fmt.Errorf("3 不支持的类型 %s", v.String()))
	}
	return ""
}

// 获取命名字段的类型字符串，如果是基础类型, 则直接返回对应的类型字符串（比如 int, uint, string, bool...）
// 如果是结构体，则是 name + "Def"
// 如果是 Map，这是 "KV" + Key + Value
func getNamedTypName(name string, typ types.Type) string {
	switch v := typ.(type) {
	case *types.Basic:
		return name
	case *types.Struct:
		return genStructName(name)
	case *types.Map:
		return fmt.Sprintf("KV%s%s", v.Key().String(), getTypString(v.Elem()))
	default:
		return name
	}
}

// 获取 attr.StrMap 或者 attr.Int32Map 的 getter 方法名
// bool 返回 当使用了 strMap.${Getter} 后， 是否需要类型转换
func getFieldAttrGetterFnName(typ types.Type) (string, bool) {
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
		return attrGetFuncName, false
	default:
		return "Value", true
	}
}
