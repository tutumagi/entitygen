package attr

import (
	"encoding/json"
	"fmt"
	"reflect"
	"strings"
	"unicode"

	dynamicstruct "github.com/ompluscator/dynamic-struct"
	"go.mongodb.org/mongo-driver/bson"
)

type Field interface {
	_InlineField
	setChangeKey(k string)
}

type _InlineField interface {
	SetParent(k string, parent Field)
}

// // 每个实体都有自己的实体ID，角色的实体ID就是角色ID
// type _Empty struct {
// 	ID string `key:"id"`
// }

type AttrFlag uint16

const (
	// 属性变化时，会通知给对应的客户端
	afClient AttrFlag = 1 << iota
	// 属性变化时，会通知给关心它的周围的客户端
	afOtherClient
	// 默认所有属性都会在 base 里面
	afBase
	// 属性变化时，会通知到对应 cell
	afCell
	// 属性变化时，会通知到对应 cell，并通知给关心它的其他实体
	afOtherCell

	// AfAllClients       = afOtherCell | afCell | afClient | afOtherClient
	// AfBase             = afBase
	// AfBaseAndClient    = afClient | afBase
	// AfCellPrivate      = afCell
	// AfCellPublic       = afCell | afOtherCell
	// AfCellPublicAndOwn = afOtherCell | afCell | afClient
	// AfOtherClients     = afOtherCell | afCell | afOtherClient
	// AfOwnClient        = afCell | afClient

	AfBase         = afBase
	AfCell         = afCell
	AfBaseAndCell  = afBase | afCell
	AfOtherClients = AfBaseAndCell | afOtherClient
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
		Uint8,
		Uint16,
		Uint32,
		Uint64,
		Float32,
		Float64,
		Bool,
		String:
		return true
	default:
		return false
	}
}

func primaryUnmarshal(att AttrTyp) func(b []byte) (interface{}, error) {
	switch att {
	case Int:
		return func(b []byte) (interface{}, error) {
			var a int
			err := json.Unmarshal(b, &a)
			return a, err
		}
	case UInt:
		return func(b []byte) (interface{}, error) {
			var a uint
			err := json.Unmarshal(b, &a)
			return a, err
		}
	case Int8:
		return func(b []byte) (interface{}, error) {
			var a int8
			err := json.Unmarshal(b, &a)
			return a, err
		}
	case Int16:
		return func(b []byte) (interface{}, error) {
			var a int16
			err := json.Unmarshal(b, &a)
			return a, err
		}
	case Int32:
		return func(b []byte) (interface{}, error) {
			var a int32
			err := json.Unmarshal(b, &a)
			return a, err
		}
	case Int64:
		return func(b []byte) (interface{}, error) {
			var a int64
			err := json.Unmarshal(b, &a)
			return a, err
		}
	case Uint8:
		return func(b []byte) (interface{}, error) {
			var a uint8
			err := json.Unmarshal(b, &a)
			return a, err
		}
	case Uint16:
		return func(b []byte) (interface{}, error) {
			var a uint16
			err := json.Unmarshal(b, &a)
			return a, err
		}
	case Uint32:
		return func(b []byte) (interface{}, error) {
			var a uint32
			err := json.Unmarshal(b, &a)
			return a, err
		}
	case Uint64:
		return func(b []byte) (interface{}, error) {
			var a uint64
			err := json.Unmarshal(b, &a)
			return a, err
		}
	case Float32:
		return func(b []byte) (interface{}, error) {
			var a float32
			err := json.Unmarshal(b, &a)
			return a, err
		}
	case Float64:
		return func(b []byte) (interface{}, error) {
			var a float64
			err := json.Unmarshal(b, &a)
			return a, err
		}
	case Bool:
		return func(b []byte) (interface{}, error) {
			var a bool
			err := json.Unmarshal(b, &a)
			return a, err
		}
	case String:
		return func(b []byte) (interface{}, error) {
			var a string
			err := json.Unmarshal(b, &a)
			return a, err
		}
	default:
		panic(fmt.Sprintf("unsupport return no primary unmarshal func. kind:%d", reflect.TypeOf(att).Kind()))
	}
}

var (
	Int    AttrTyp = int(0)
	UInt   AttrTyp = uint(0)
	Int8   AttrTyp = int8(0)
	Int16  AttrTyp = int16(0)
	Int32  AttrTyp = int32(0)
	Int64  AttrTyp = int64(0)
	Uint8  AttrTyp = uint8(0)
	Uint16 AttrTyp = uint16(0)
	Uint32 AttrTyp = uint32(0)
	Uint64 AttrTyp = uint64(0)

	Float32 AttrTyp = float32(0)
	Float64 AttrTyp = float64(0)

	Bool AttrTyp = bool(false)

	String AttrTyp = string("")
)

type FieldDef struct {
	// 该字段的 flag
	flag AttrFlag
	// 是否需要存储到 db 里面
	storeDB bool
	// 是否是基础类型（比如整型，浮点型，bool，string）
	primary bool
	// 反射用的
	typp reflect.Type
	// 值，动态构建 struct 需要用的
	typv interface{}

	primaryUnmarshal func(b []byte) (interface{}, error)
}

func (f *FieldDef) Flag() AttrFlag {
	return f.flag
}

func (f *FieldDef) HasCell() bool {
	return f.flag&afCell > 0
}

func (f *FieldDef) HasClient() bool {
	return f.flag&afClient > 0
}

func (f *FieldDef) HasOtherCell() bool {
	return f.flag&afOtherCell > 0
}

func (f *FieldDef) HasOtherClient() bool {
	return f.flag&afOtherClient > 0
}

func (f *FieldDef) StoreDB() bool {
	return f.storeDB
}

func (f *FieldDef) IsPrimary() bool {
	return f.primary
}

func (f *FieldDef) V() interface{} {
	return f.typv
}

func (f *FieldDef) P() reflect.Type {
	return f.typp
}

func (f *FieldDef) UnmarshalPrimary(b []byte) (interface{}, error) {
	if f.primary {
		return f.primaryUnmarshal(b)
	}
	panic(fmt.Sprintf("unmarshal a not primary field. kind:%d", f.typp.Kind()))
}

type Meta struct {
	fields map[string]*FieldDef

	dynStruct dynamicstruct.DynamicStruct

	creater     func() interface{}
	createSlice func() interface{}
}

func NewMeta(creater func() interface{}, createSlice func() interface{}) *Meta {
	m := &Meta{}
	m.creater = creater
	m.createSlice = createSlice
	return m
}

func (meta *Meta) DefAttr(key string, typ AttrTyp, flag AttrFlag, storeDB bool) {
	if meta.fields == nil {
		meta.fields = make(map[string]*FieldDef, 10)
	}
	f := &FieldDef{
		flag:    flag,
		typv:    typ,
		typp:    reflect.TypeOf(typ),
		storeDB: storeDB,
		primary: isPrimary(typ),
	}
	meta.fields[key] = f
	if f.primary {
		f.primaryUnmarshal = primaryUnmarshal(typ)
	}
}

func (meta *Meta) GetDef(key string) *FieldDef {
	return meta.fields[key]
}

func (meta *Meta) Create() interface{} {
	return meta.creater()
}

func (meta *Meta) CreateSlice() interface{} {
	return meta.createSlice()
}

func (meta *Meta) DynamicStruct() interface{} {
	return meta.builder().New()
}

func (meta *Meta) DynamicSliceOfStruct() interface{} {
	// TODO 有没有可能构造Slice时 加cap
	return meta.builder().NewSliceOfStructs()
}

func (meta *Meta) builder() dynamicstruct.DynamicStruct {
	if meta.dynStruct == nil {
		// builder := dynamicstruct.ExtendStruct(_Empty{}) // 这个是默认数据结构中都有一个 ID("id")
		builder := dynamicstruct.NewStruct()
		// builder.AddField("ID", "", `json:"id" bson:"id"`)
		for k, v := range meta.fields {
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
				fmt.Sprintf(`json:"%s, omitempty" bson:"%s"`, k, tagStr),
			)
		}

		meta.dynStruct = builder.Build()
	}
	return meta.dynStruct
}

// 通过 dynamicStruct 解析到的struct，转为 map[string]interface{}
func (meta *Meta) Unmarshal(srcStruct interface{}, sm *StrMap) *StrMap {
	return meta.readerToMap(dynamicstruct.NewReader(srcStruct), sm)
}

// 通过 dynamicStruct 解析到的struct，转为 []map[string]interface{}
func (meta *Meta) UnmarshalSlice(srcStruct interface{}) []*StrMap {
	// srcStruct 类型为 *[]*StrMap
	// v := reflect.ValueOf(srcStruct)
	// if v.Kind() == reflect.Ptr {
	// 	v = v.Elem()
	// }
	// if v.Kind() == reflect.Slice {

	// }
	var attrs = []*StrMap{}
	readers := dynamicstruct.NewReader(srcStruct).ToSliceOfReaders()
	for _, r := range readers {
		attrs = append(attrs, meta.readerToMap(r, nil))
	}

	return attrs
}

// 将 dynamicstruct.Reader 转为 map[string]interface{}
func (meta *Meta) readerToMap(r dynamicstruct.Reader, attrs *StrMap) *StrMap {
	if attrs == nil {
		attrs = NewStrMap(nil)
	}
	for _, field := range r.GetAllFields() {
		name := lowerFirst(field.Name())
		v := field.Interface()
		attrs.Set(name, v)

		if !meta.GetDef(name).IsPrimary() {
			v.(_InlineField).SetParent(name, attrs)
		}
	}

	return attrs
}

func (meta *Meta) UnmarshalBson(bytes []byte, sm *StrMap) (*StrMap, error) {
	dynStruct := meta.DynamicStruct()
	err := bson.Unmarshal(bytes, dynStruct)
	if err != nil {
		return nil, err
	}
	return meta.Unmarshal(dynStruct, sm), nil
}

func (meta *Meta) UnmarshalJson(bytes []byte, sm *StrMap) (*StrMap, error) {
	dynStruct := meta.DynamicStruct()
	err := json.Unmarshal(bytes, dynStruct)
	if err != nil {
		return nil, err
	}
	return meta.Unmarshal(dynStruct, sm), nil
}

func lowerFirst(s string) string {
	for _, c := range s {
		return string(unicode.ToLower(c)) + s[1:]
	}
	return s
}

type IAttr interface {
	Undertype() interface{}
}

type A map[string]interface{}
