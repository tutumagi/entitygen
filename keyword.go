package main

import (
	"fmt"
	"go/types"
	"strings"

	. "github.com/dave/jennifer/jen"
	"gitlab.nftgaga.com/usm/game/entitygen/attr"
)

const (
	thisKeyword       = "a"
	attrPackageName   = "gitlab.nftgaga.com/usm/game/entitygen/attr"
	setParentFuncName = "SetParent"

	// 实体内置的属性 id
	buildinIDKey    = "eid"
	buildinIDSetter = "SetId"
	buildinIDGetter = "GetId"

	buildinPosKey = "pos"
	buildinRotKey = "rot"

	// 更新对象
	// updateFuncName = "update"
)

// 一些预设的类型或者关键字
var (
	// attr.Field
	attrField = func() *Statement { return Qual(attrPackageName, "Field") }
	// attr.Meta
	attrMeta = func() *Statement { return Qual(attrPackageName, "Meta") }

	// attr.NewMeta
	attrNewMeta = func() *Statement { return Qual(attrPackageName, "NewMeta") }

	attrVec3 = func() *Statement { return Qual(attrPackageName, "Vec3") }
)

// 根据 要生成的类型名字，和依赖的 attr 里面的名字生成一些预设的 statement
func aboutThisCode(
	structName string,
	attrTypeName string,
) (
	attrType func() *Statement,
	thisFn func() *Statement,
	convertThis func() *Statement,
	convertAttrType func(string) *Statement,
) {
	// 比如 attr.Int32Map attr.StrMap attr.Slice
	attrType = func() *Statement { return Qual(attrPackageName, attrTypeName) }
	// 比如 a *XXXX （主要用在类方法时的 this 定义）
	thisFn = func() *Statement { return Id(thisKeyword).Op("*").Id(structName) }
	// 比如 (*attr.Int32Map)(a)
	convertThis = func() *Statement { return convertAttrType(thisKeyword) }
	// 比如 (*attr.Int32Map)(name)
	convertAttrType = func(name string) *Statement { return Parens(Id("*").Add(attrType())).Parens(Id(name)) }
	return
}

func EmptyCtor(typName string) string {
	return fmt.Sprintf("Empty%s", typName)
}

func NormalCtor(typName string) string {
	return fmt.Sprintf("New%s", typName)
}

func CopyCtor(typName string) string {
	return fmt.Sprintf("Copy%s", typName)
}

// 通过原始 struct 名字，获取生成的 struct 名字
func StructTypeName(sourceTypeName string) string {
	return sourceTypeName
}

// 通过原始 struct 名字，获取生成的 struct 的 meta 变量名字
func StructMetaName(srcName string, isEntity bool) string {
	// return strings.ToLower(srcName) + "Meta"
	result := srcName + "Meta"
	if !isEntity {
		return attr.LowerFirst(result)
	}
	return result
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
