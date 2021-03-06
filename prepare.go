package main

import (
	"fmt"
	"go/types"
	"reflect"
	"strings"
	"unicode"

	. "github.com/dave/jennifer/jen"
)

// 标明 rot 和 pos 的一些标记，只给 entStructInfo 使用
type buildinVec3Info struct {
	hasCell   bool
	hasClient bool
	storedb   bool
}

// 记录该实体是否需要生成 rot 和 pos 的信息（以及他们的 flag：是否存 db，是否需要到 cellapp 里面去）
type entStructInfo struct {
	rot *buildinVec3Info // 为空，则不生成内置的 rot 字段；不为空，则表示需要生成内置的 rot 字段
	pos *buildinVec3Info // 为空，则不生成内置的 pos 字段；不为空，则表示需要生成内置的 pos 字段
}

type fieldInfo struct {
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
	// 自定义类型目前也是自定义类型
	typName string
	// zero value 对应的 Code
	emptyValue Code
}

func getEntStruct(structType *types.Struct) *entStructInfo {
	entInfo := &entStructInfo{}
	for i := 0; i < structType.NumFields(); i++ {
		field := structType.Field(i)
		name := field.Name()

		tagValue := reflect.StructTag(structType.Tag(i))
		storeDB, _, flagCell, client := checkFieldTagValue(name, tagValue)

		lowerName := strings.ToLower(name)
		// 检查属性定义里面有没有跟内置的属性重复的地方
		if strings.ToLower(buildinIDKey) == lowerName {
			panic(fmt.Sprintf("%s 是内置字段，不能使用\n", buildinIDKey))
		}
		if strings.ToLower(buildinPosKey) == lowerName {
			fmt.Printf("需要生成 pos 字段\n")
			entInfo.pos = &buildinVec3Info{
				hasCell:   flagCell,
				hasClient: client,
				storedb:   storeDB,
			}
		}
		if strings.ToLower(buildinRotKey) == lowerName {
			fmt.Printf("需要生成 rot 字段\n")
			entInfo.rot = &buildinVec3Info{
				hasCell:   flagCell,
				hasClient: client,
				storedb:   storeDB,
			}
		}
	}
	return entInfo
}

func getStructFields(structType *types.Struct, isEntityDef bool) ([]*fieldInfo, *entStructInfo) {
	result := make([]*fieldInfo, 0, structType.NumFields())
	for i := 0; i < structType.NumFields(); i++ {
		field := structType.Field(i)
		name := field.Name()

		typ := field.Type()
		storeDB := true
		flagBase := true // 目前的实现里面属性肯定会存储在 base 里面
		flagCell := false
		client := true

		tagValue := reflect.StructTag(structType.Tag(i))
		key, ok := tagValue.Lookup("key")
		if !ok {
			// 如果 tag 没有 key，则使用 fieldName，并把第一个字母小写作为 key
			// 比如 fieldName 为 Desk，则 key 为 desk
			// 对应的 Getter Setter 方法名为 GetDesk, SetDesk
			key = func() string {
				for _, c := range name {
					return string(unicode.ToLower(c)) + name[1:]
				}
				return name
			}()
		}
		getterSetterName := strings.Title(key)

		// 如果是实体定义 才需要检查 flag, storedb 和 client 的标签
		if isEntityDef {
			lowerName := strings.ToLower(name)
			// 检查属性定义里面有没有跟内置的属性重复的地方
			if strings.ToLower(buildinIDKey) == lowerName {
				panic(fmt.Sprintf("%s 是内置字段，不能使用", buildinIDKey))
			}
			if strings.ToLower(buildinPosKey) == lowerName {
				continue
			}
			if strings.ToLower(buildinRotKey) == lowerName {
				continue
			}
			storeDB, flagBase, flagCell, client = checkFieldTagValue(name, tagValue)
		}

		typName := getTypString(typ)
		attrGetter, _ := getFieldAttrGetterFnName(typ)
		result = append(result, &fieldInfo{
			name:       name,
			key:        key,
			typ:        typ,
			storeDB:    storeDB,
			base:       flagBase,
			cell:       flagCell,
			client:     client,
			emptyValue: getEmptyValue(typName, typ),
			getter:     Id(fmt.Sprintf("Get%s", getterSetterName)),
			setter:     Id(fmt.Sprintf("Set%s", getterSetterName)),
			setParam:   Id(key).Id(typName),
			typName:    typName,
			attrGetter: attrGetter,
		})
	}

	if isEntityDef {
		return result, getEntStruct(structType)
	}
	return result, nil
}

// 获取类型的空值
func getEmptyValue(typName string, typ types.Type) Code {
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
		// 如果这里的空值，返回的仍然是空值的构造方法的话，如果某个自定义类型的字段也是该类型，就会无限递归构造，所以空值构造器，自定义类型使用 nil 值
		// 这样的话在 attr 中 SetParent 和 ToMap 这些方法中要对 `this` 是否为空做判断
	// case *types.Map:
	// 	return Id(EmptyCtor(trimHeadStar(typName))).Call()
	// case *types.Struct:
	// 	return Id(EmptyCtor(trimHeadStar(typName))).Call()
	// case *types.Named:
	// 	return Id(EmptyCtor(trimHeadStar(typName))).Call()
	// case *types.Slice:
	// 	return Id(EmptyCtor(trimHeadStar(typName))).Call()
	// case *types.Pointer:
	// 	return getEmptyValue(trimHeadStar(typName), v.Elem())
	// default:
	// 	failErr(fmt.Errorf("空值 Code 获取失败, 不支持的 type:%s", typ))
	default:
		return Nil()
	}
	// return Id("")
}

// // 获取 nil 值，基础类型就是和空值一样，非基础类型就是 nil
// func getNilValue(typ types.Type) Code {
// 	switch v := typ.(type) {
// 	case *types.Basic:
// 		switch v.Kind() {
// 		case types.String, types.UntypedString:
// 			return Lit("")
// 		case types.Int, types.Uint, types.Int8, types.Uint8, types.Int16, types.Uint16, types.Int32, types.Uint32, types.Int64, types.Uint64, types.Float32, types.Float64:
// 			return Lit(0)
// 		case types.Bool:
// 			return Lit(false)
// 		default:
// 			return Nil()
// 		}
// 	default:
// 		return Nil()
// 	}
// }

func getTypString(typ types.Type) string {
	// 获取命名字段的类型字符串，如果是基础类型, 则直接返回对应的类型字符串（比如 int, uint, string, bool...）
	// 如果是结构体，则是 name + "Def"
	// 如果是 Map，这是 "KV" + Key + Value
	getNamedTypName := func(name string, typ types.Type) string {
		switch v := typ.(type) {
		case *types.Basic:
			return name
		case *types.Struct:
			return StructTypeName(name)
		case *types.Map:
			return MapTypeName(v)
		default:
			return name
		}
	}

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
		return fmt.Sprintf("*%s", MapTypeName(v))
	case *types.Slice:
		return fmt.Sprintf("*%s", SliceTypeName(v))
	default:
		failErr(fmt.Errorf("不支持的类型 %s", v.String()))
	}
	return ""
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

func checkFieldTagValue(fieldName string, tagValue reflect.StructTag) (storeDB bool, flagBase bool, flagCell bool, client bool) {
	storeDB = true
	flagBase = true // 目前的实现里面属性肯定会存储在 base 里面
	flagCell = false
	client = true

	{
		storeDBStr, ok := tagValue.Lookup("storedb")
		if !ok {
			failErr(fmt.Errorf("field:%s 必须有tag:storedb", fieldName))
		}
		if storeDBStr != "true" && storeDBStr != "false" {
			failErr(fmt.Errorf("field:%s storedb(%s) 必须是 true 或者 false", fieldName, storeDBStr))
		}

		storeDB = storeDBStr == "true"
	}

	{
		clientStr, ok := tagValue.Lookup("client")
		if !ok {
			failErr(fmt.Errorf("field:%s 必须有tag:client", fieldName))
		}
		if clientStr != "true" && clientStr != "false" {
			failErr(fmt.Errorf("field:%s client(%s) 必须是 true 或者 false", fieldName, clientStr))
		}

		client = clientStr == "true"
	}

	{
		flagStr, ok := tagValue.Lookup("flag")
		if !ok {
			failErr(fmt.Errorf("field:%s 必须有tag:flag", fieldName))
		}
		if flagStr != "base" && flagStr != "cell" {
			failErr(fmt.Errorf("field:%s flag(%s) 必须是 base 或者 cell", fieldName, flagStr))
		}

		flagCell = flagStr == "cell"
	}

	return
}
