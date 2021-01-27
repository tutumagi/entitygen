package attr

import (
	"encoding/json"
	"fmt"
	"reflect"
	"strings"

	dynamicstruct "github.com/ompluscator/dynamic-struct"
	"go.mongodb.org/mongo-driver/bson"
)

type Field interface {
	setChangeKey(k string)
}

// 每个实体都有自己的实体ID，角色的实体ID就是角色ID
type _Empty struct {
	ID string `bson:"id" json:"id"`
}

type AttrFlag uint16

const (
	afClient AttrFlag = 1 << iota
	afOtherClient
	afBase
	afCell
	afOtherCell

	// AfAllClients       = afOtherCell | afCell | afClient | afOtherClient
	// AfBase             = afBase
	// AfBaseAndClient    = afClient | afBase
	// AfCellPrivate      = afCell
	// AfCellPublic       = afCell | afOtherCell
	// AfCellPublicAndOwn = afOtherCell | afCell | afClient
	// AfOtherClients     = afOtherCell | afCell | afOtherClient
	// AfOwnClient        = afCell | afClient

	AfBase        = afBase
	AfCell        = afCell
	AfBaseAndCell = afBase | afCell
)

type AttrTyp interface{}

func isPrimary(a AttrTyp) bool {
	switch a {
	case Int,
		UInt,
		Int8,
		Int16,
		Int32,
		Int64,
		UInt8,
		UInt16,
		UInt32,
		UInt64,
		Float32,
		Float64,
		Bool,
		String:
		return true
	default:
		return false
	}
}

var (
	Int    AttrTyp = int(0)
	UInt   AttrTyp = uint(0)
	Int8   AttrTyp = int8(0)
	Int16  AttrTyp = int16(0)
	Int32  AttrTyp = int32(0)
	Int64  AttrTyp = int64(0)
	UInt8  AttrTyp = uint8(0)
	UInt16 AttrTyp = uint16(0)
	UInt32 AttrTyp = uint32(0)
	UInt64 AttrTyp = uint64(0)
	// 这样做的原因，会导致内存增长，
	// 但对业务方来说，使用起来不容易出bug，
	// 否则，在定义属性，写属性，读属性时 都必须保持类型一致，否则就会导致同一个属性key，拿到的值不一致的bug
	// IntAttr    AttrTyp = float64(0)
	// UIntAttr   AttrTyp = float64(0)
	// Int8Attr   AttrTyp = float64(0)
	// Int16Attr  AttrTyp = float64(0)
	// Int32Attr  AttrTyp = float64(0)
	// Int64Attr  AttrTyp = float64(0)
	// UInt8Attr  AttrTyp = float64(0)
	// UInt16Attr AttrTyp = float64(0)
	// UInt32Attr AttrTyp = float64(0)
	// UInt64Attr AttrTyp = float64(0)

	Float32 AttrTyp = float32(0)
	Float64 AttrTyp = float64(0)

	Bool AttrTyp = bool(false)

	String AttrTyp = string("")

	MapStrStrAttr   AttrTyp = map[string]string{}
	MapStrInt32Attr AttrTyp = map[string]int32{}

	MapInt32StrAttr   AttrTyp = map[int32]string{}
	MapInt32Int32Attr AttrTyp = map[int32]int32{}

	SliceStrAttr   AttrTyp = []string{}
	SliceInt32Attr AttrTyp = []int32{}
	// 如果不是基础类型，则自己传入 值类型
	// InterfaceAttr AttrTyp =
)

type FieldDef struct {
	flag AttrFlag
	// 是否需要存储到 db 里面
	storeDB bool
	// 是否是基础类型（比如整型，浮点型，bool，string）
	primary bool
	// 反射用的
	typp reflect.Type
	// 值，动态构建 struct 需要用的
	typv interface{}
}

func (f *FieldDef) Flag() AttrFlag {
	return f.flag
}

func (f *FieldDef) StoreDB() bool {
	return f.storeDB
}

func (f *FieldDef) IsPrimary() bool {
	return f.primary
}

type Def struct {
	attrDefs map[string]*FieldDef

	dynStruct dynamicstruct.DynamicStruct
}

func (desc *Def) DefAttr(key string, typ AttrTyp, flag AttrFlag, storeDB bool) {
	if desc.attrDefs == nil {
		desc.attrDefs = make(map[string]*FieldDef, 10)
	}
	desc.attrDefs[key] = &FieldDef{
		flag:    flag,
		typv:    typ,
		typp:    reflect.TypeOf(typ),
		storeDB: storeDB,
		primary: isPrimary(typ),
	}
}

func (desc *Def) GetDef(key string) *FieldDef {
	return desc.attrDefs[key]
}

func (desc *Def) DynamicStruct() interface{} {
	return desc.builder().New()
}

func (desc *Def) DynamicSliceOfStruct() interface{} {
	// TODO 有没有可能构造Slice时 加cap
	return desc.builder().NewSliceOfStructs()
}

func (desc *Def) builder() dynamicstruct.DynamicStruct {
	if desc.dynStruct == nil {
		// builder := dynamicstruct.ExtendStruct(_Empty{})	// 这个是默认数据结构中都有一个 ID("id")
		builder := dynamicstruct.NewStruct()
		for k, v := range desc.attrDefs {
			tagStr := "-"
			if v.storeDB {
				tagStr = k
			}
			// Field的 name 必须是大写开头的，因为go语言 反射必须是外部包可见的field
			// 写到db是使用的 bson， json是内存中 marshal unmarshal使用的，所以json不忽略，
			// 当不需要存储到db时，bson 忽略，使用 `-` tag
			builder.AddField(
				strings.Title(k), // 首字母大写
				v.typv,
				fmt.Sprintf(`json:"%s" bson:"%s"`, k, tagStr),
			)
		}

		desc.dynStruct = builder.Build()
	}
	return desc.dynStruct
}

// 通过 dynamicStruct 解析到的struct，转为 map[string]interface{}
func (desc *Def) unmarshal(srcStruct interface{}) map[string]interface{} {
	return desc.readerToMap(dynamicstruct.NewReader(srcStruct))
}

// 通过 dynamicStruct 解析到的struct，转为 map[string]interface{}
func (desc *Def) unmarshalSlice(srcStruct interface{}) []map[string]interface{} {
	var attrs = []map[string]interface{}{}
	readers := dynamicstruct.NewReader(srcStruct).ToSliceOfReaders()
	for _, r := range readers {
		attrs = append(attrs, desc.readerToMap(r))
	}

	return attrs
}

// 将 dynamicstruct.Reader 转为 map[string]interface{}
func (desc *Def) readerToMap(r dynamicstruct.Reader) map[string]interface{} {
	var attrs = map[string]interface{}{}
	for _, field := range r.GetAllFields() {
		name := strings.ToLower(field.Name()) // TODO 这里有性能瓶颈，可以考虑 修改dynamicstruct 的源码，去缓存这个 小写开头的字符串
		attrs[name] = field.Interface()
	}

	return attrs
}

func (desc *Def) UnmarshalBson(bytes []byte) (map[string]interface{}, error) {
	dynStruct := desc.DynamicStruct()
	err := bson.Unmarshal(bytes, dynStruct)
	if err != nil {
		return nil, err
	}
	return desc.unmarshal(dynStruct), nil
}

func (desc *Def) UnmarshalJson(bytes []byte) (map[string]interface{}, error) {
	dynStruct := desc.DynamicStruct()
	err := json.Unmarshal(bytes, dynStruct)
	if err != nil {
		return nil, err
	}
	return desc.unmarshal(dynStruct), nil
}
