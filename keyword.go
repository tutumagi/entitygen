package main

import (
	"fmt"
	"go/types"
	"strings"
)

const (
	thisKeyword = "a"
)

var (

// 	zeroTypeCtor = func(v types.Type) Code {

// 	}
)

func EmptyCtor(typName string) string {
	return fmt.Sprintf("Empty%s", typName)
}

func NormalCtor(typName string) string {
	return fmt.Sprintf("New%s", typName)
}

// 通过原始 struct 名字，获取生成的 struct 名字
func StructTypeName(sourceTypeName string) string {
	return sourceTypeName + "Def"
}

// 通过原始 struct 名字，获取生成的 struct 的 meta 变量名字
func StructMetaName(srcName string) string {
	return strings.ToLower(srcName) + "Meta"
}

// 生成的 Map 名字 KV{Key}{Val}
func MapTypeName(v *types.Map) string {
	key := strings.Title(v.Key().String())
	if key == "String" {
		key = "Str"
	}

	val := trimHeadStar(getTypString(v.Elem()))

	val = strings.Title(val)
	if val == "String" {
		val = "Str"
	}
	return fmt.Sprintf("KV%s%s", key, val)
}

func SliceTypeName(v *types.Slice) string {
	val := trimHeadStar(getTypString(v.Elem()))

	val = strings.Title(val)
	if val == "String" {
		val = "Str"
	}

	return fmt.Sprintf("%sSlice", val)
}

// 如果是 *Desk, 则返回 Desk
func trimHeadStar(str string) string {
	if strings.HasPrefix(str, "*") {
		str = strings.TrimLeft(str, "*")
	}
	return str
}
