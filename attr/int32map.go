package attr

import (
	"fmt"
	"strings"
	"sync"
)

// Int32Map map attr
type Int32Map struct {
	// 当前这个数据 在根结构中的 key 名字
	// 比如 {
	// 	"a": {"b": {"c": 123 } }
	// }
	// 那么 b，c 的 parentKey 都是 a
	parentKey string
	data      map[int32]interface{}

	// 存这个数据的父节点
	parent Field
}

var int32MapPool *sync.Pool = &sync.Pool{
	New: func() interface{} {
		return &Int32Map{
			parentKey: "",
			data:      map[int32]interface{}{},

			parent: nil,
		}
	},
}

func NewInt32Map(data map[int32]interface{}) *Int32Map {
	int32map := int32MapPool.Get().(*Int32Map)
	int32map.parentKey = ""
	int32map.parent = nil
	int32map.data = data
	return int32map
}

func ReleaseInt32Map(mm *Int32Map) {
	mm.data = map[int32]interface{}{}
	mm.parentKey = ""
	mm.parent = nil
	int32MapPool.Put(mm)
}

func (a *Int32Map) String() string {
	var sb strings.Builder
	sb.WriteString("MapInt32Attr{")
	isFirstField := true
	for k, v := range a.data {
		if !isFirstField {
			sb.WriteString(", ")
		}

		fmt.Fprintf(&sb, "%#v", k)
		sb.WriteString(": ")
		switch a := v.(type) {
		case *Int32Map:
			sb.WriteString(a.String())
		// case *AttrList:
		// 	sb.WriteString(a.String())
		default:
			fmt.Fprintf(&sb, "%#v", v)
		}
		isFirstField = false
	}
	sb.WriteString("}")
	return sb.String()
}

func (a *Int32Map) ToMap() map[int32]interface{} {
	if a == nil {
		return nil
	}
	result := map[int32]interface{}{}
	for k, v := range a.data {
		if !isNil(v) {
			result[k] = v
		}
	}
	return result
	// return a.data
}

func (a *Int32Map) Has(k int32) bool {
	_, ok := a.data[k]
	return ok
}

func (a *Int32Map) ForEach(fn func(k int32, v interface{}) bool) {
	for k, v := range a.data {
		if !fn(k, v) {
			break
		}
	}
}

func (a *Int32Map) Delete(key int32) bool {
	if _, ok := a.data[key]; ok {
		delete(a.data, key)
		a.change()
		return true
	}
	return false
}

func (a *Int32Map) FastDelete(key int32) {
	delete(a.data, key)
	a.change()
}

func (a *Int32Map) change() {
	if a.parent != nil {
		a.parent.setChangeKey(a.parentKey)
	}
}

func (a *Int32Map) setChangeKey(k string) {
	a.change()
}

func (a *Int32Map) SetParent(k string, parent Field) {
	if a == nil {
		return
	}
	if (a.parentKey != "" && a.parentKey != k) || (a.parent != nil && a.parent != parent) {
		panic(
			fmt.Sprintf(
				"has already set parent oldKey:%s newKey:%s oldParent:%s newParent:%s",
				a.parentKey, k,
				a.parent, parent,
			),
		)
	}
	a.parentKey = k
	a.parent = parent
}

func (a *Int32Map) SetData(data map[int32]interface{}) {
	a.data = data
}

func (a *Int32Map) Set(key int32, val interface{}) {
	a.data[key] = val
	// 这里缓存 修改的 key
	// 还有一种做法是 改变立马通知除去
	a.change()
}

// Bool returns value with Bool type
func (a *Int32Map) Bool(key int32) bool {
	v, ok := a.data[key]
	if !ok {
		return false
	}
	r, ok := v.(bool)
	if !ok {
		return false
	}
	return r
}

// String returns value with String type
func (a *Int32Map) Str(key int32) string {
	v, ok := a.data[key]
	if !ok {
		return ""
	}
	r, ok := v.(string)
	if !ok {
		return ""
	}
	return r
}

// Value returns value with interface{} type
func (a *Int32Map) Value(key int32) interface{} {
	v, ok := a.data[key]
	if !ok {
		return nil
	}
	return v
}

// Int returns value with Int type
func (a *Int32Map) Int(key int32) int {
	v, ok := a.data[key]
	if !ok {
		return 0
	}
	r, ok := v.(int)
	if !ok {
		return 0
	}
	return r
}

// Int8 returns value with Int8 type
func (a *Int32Map) Int8(key int32) int8 {
	v, ok := a.data[key]
	if !ok {
		return 0
	}
	r, ok := v.(int8)
	if !ok {
		return 0
	}
	return r
}

// Int16 returns value with Int16 type
func (a *Int32Map) Int16(key int32) int16 {
	v, ok := a.data[key]
	if !ok {
		return 0
	}
	r, ok := v.(int16)
	if !ok {
		return 0
	}
	return r
}

// Int32 returns value with Int32 type
func (a *Int32Map) Int32(key int32) int32 {
	v, ok := a.data[key]
	if !ok {
		return 0
	}
	r, ok := v.(int32)
	if !ok {
		return 0
	}
	return r
}

// Int64 returns value with Int64 type
func (a *Int32Map) Int64(key int32) int64 {
	v, ok := a.data[key]
	if !ok {
		return 0
	}
	r, ok := v.(int64)
	if !ok {
		return 0
	}
	return r
}

// UInt returns value with UInt type
func (a *Int32Map) UInt(key int32) uint {
	v, ok := a.data[key]
	if !ok {
		return 0
	}
	r, ok := v.(uint)
	if !ok {
		return 0
	}
	return r
}

// Uint8 returns value with Uint8 type
func (a *Int32Map) Uint8(key int32) uint8 {
	v, ok := a.data[key]
	if !ok {
		return 0
	}
	r, ok := v.(uint8)
	if !ok {
		return 0
	}
	return r
}

// Uint16 returns value with Uint16 type
func (a *Int32Map) Uint16(key int32) uint16 {
	v, ok := a.data[key]
	if !ok {
		return 0
	}
	r, ok := v.(uint16)
	if !ok {
		return 0
	}
	return r
}

// Uint32 returns value with Uint32 type
func (a *Int32Map) Uint32(key int32) uint32 {
	v, ok := a.data[key]
	if !ok {
		return 0
	}
	r, ok := v.(uint32)
	if !ok {
		return 0
	}
	return r
}

// Uint64 returns value with Uint64 type
func (a *Int32Map) Uint64(key int32) uint64 {
	v, ok := a.data[key]
	if !ok {
		return 0
	}
	r, ok := v.(uint64)
	if !ok {
		return 0
	}

	return r
}

// Float32 returns value with float32 type
func (a *Int32Map) Float32(key int32) float32 {
	v, ok := a.data[key]
	if !ok {
		return 0
	}
	r, ok := v.(float32)
	if !ok {
		return 0
	}
	return r
}

// Float64 returns value with float64 type
func (a *Int32Map) Float64(key int32) float64 {
	v, ok := a.data[key]
	if !ok {
		return 0
	}
	r, ok := v.(float64)
	if !ok {
		return 0
	}

	return r
}

func (a *Int32Map) Equal(other *Int32Map) bool {
	if a == nil && other == nil {
		return true
	}
	if (a == nil && other != nil) || (a != nil && other == nil) {
		return false
	}
	if len(a.data) != len(other.data) {
		return false
	}
	equal := true
	for k, v := range a.data {
		if uu, ok := v.(IAttr); ok {
			if im, ok := uu.Undertype().(*Int32Map); ok {
				otherV := other.Value(k)
				if otherV != nil {
					if otherVV, ok := otherV.(IAttr); ok {
						if othervvv, ok := otherVV.Undertype().(*Int32Map); ok {
							if im.Equal(othervvv) {
								continue
							}
						}
					}
				}
				equal = false
				break
			}
			if im, ok := uu.Undertype().(*StrMap); ok {
				otherV := other.Value(k)
				if otherV != nil {
					if otherVV, ok := otherV.(IAttr); ok {
						if othervvv, ok := otherVV.Undertype().(*StrMap); ok {
							if im.Equal(othervvv) {
								continue
							}
						}
					}
				}
				equal = false
				break
			}
			if im, ok := uu.Undertype().(*Slice); ok {
				otherV := other.Value(k)
				if otherV != nil {
					if otherVV, ok := otherV.(IAttr); ok {
						if othervvv, ok := otherVV.Undertype().(*Slice); ok {
							if im.Equal(othervvv) {
								continue
							}
						}
					}
				}
				equal = false
				break
			}
		}

		if v == other.Value(k) {
			continue
		} else {
			equal = false
			break
		}
	}

	return equal
}

// 返回值只读，由外部自己保证不要去改这里面的东西
func (a *Int32Map) Data() map[int32]interface{} {
	return a.data
}

func (a *Int32Map) Undertype() interface{} {
	return a
}

func (a *Int32Map) Len() int {
	return len(a.data)
}
