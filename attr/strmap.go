package attr

import (
	"fmt"
	"strings"
	"sync"
)

// StrMap str 作为 key 的 map
type StrMap struct {
	key       interface{} // 在父节点中的 key，可能是 string，可能是 int32
	parentKey string
	// 存这个数据的祖宗
	parent AttrField

	data map[string]interface{}

	changedkey map[string]struct{}
}

var strMapPool *sync.Pool = &sync.Pool{
	New: func() interface{} {
		return &StrMap{
			key:        "",
			parentKey:  "",
			parent:     nil,
			data:       map[string]interface{}{},
			changedkey: map[string]struct{}{},
		}
	},
}

func NewStrMap(data map[string]interface{}) *StrMap {
	a := strMapPool.Get().(*StrMap)
	a.key = ""
	a.parentKey = ""
	a.parent = nil

	for k, v := range data {
		a.data[k] = v
	}

	return a
}

func ReleaseStrMap(strMap *StrMap) {
	for k := range strMap.data {
		delete(strMap.data, k)
	}
	for k := range strMap.changedkey {
		delete(strMap.changedkey, k)
	}
	strMapPool.Put(strMap)
}

func (a *StrMap) SetParent(k string, parent AttrField) {
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

// func (a *StrMap) ToMap(filter ...func(k string) bool) map[string]interface{} {
func (a *StrMap) ToMap() map[string]interface{} {
	// result := map[string]interface{}{}
	// var f func(k string) bool = nil
	// if len(filter) > 0 {
	// 	f = filter[0]
	// }
	// for k, v := range a.data {
	// 	if f != nil {
	// 		if f(k) {
	// 			result[k] = v
	// 		}
	// 	} else {
	// 		result[k] = v
	// 	}
	// }
	// return result
	return a.data
}

func (a *StrMap) ForEach(fn func(k string, v interface{}) bool) {
	for k, v := range a.data {
		if !fn(k, v) {
			break
		}
	}
}

func (a *StrMap) Delete(key string) bool {
	if _, ok := a.data[key]; ok {
		delete(a.data, key)
		a.setChangeKey(key)
		return true
	}
	return false
}

func (a *StrMap) FastDelete(key string) {
	delete(a.data, key)
	a.setChangeKey(key)
}

func (a *StrMap) SetData(data map[string]interface{}) {
	a.data = data
}

func (a *StrMap) Set(key string, val interface{}) {
	a.data[key] = val
	a.setChangeKey(key)
}

// Bool returns value with Bool type
func (a *StrMap) Bool(key string) bool {
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
func (a *StrMap) Str(key string) string {
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
func (a *StrMap) Value(key string) interface{} {
	v, ok := a.data[key]
	if !ok {
		return nil
	}
	return v
}

// Int returns value with Int type
func (a *StrMap) Int(key string) int {
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
func (a *StrMap) Int8(key string) int8 {
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
func (a *StrMap) Int16(key string) int16 {
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
func (a *StrMap) Int32(key string) int32 {
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
func (a *StrMap) Int64(key string) int64 {
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
func (a *StrMap) UInt(key string) uint {
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

// UInt8 returns value with UInt8 type
func (a *StrMap) UInt8(key string) uint8 {
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

// UInt16 returns value with UInt16 type
func (a *StrMap) UInt16(key string) uint16 {
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

// UInt32 returns value with UInt32 type
func (a *StrMap) UInt32(key string) uint32 {
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

// UInt64 returns value with UInt64 type
func (a *StrMap) UInt64(key string) uint64 {
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
func (a *StrMap) Float32(key string) float32 {
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
func (a *StrMap) Float64(key string) float64 {
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

func (a *StrMap) setChangeKey(key string) {
	if a.parent == nil {
		a.changedkey[key] = struct{}{}
	} else {
		a.parent.setChangeKey(a.parentKey)
	}
}

func (a *StrMap) HasChange() bool {
	return len(a.changedkey) > 0
}

func (a *StrMap) ClearChangeKey() {
	for k := range a.changedkey {
		delete(a.changedkey, k)
	}
}

func (a *StrMap) ChangeKey() map[string]struct{} {
	return a.changedkey
}

func (a *StrMap) String() string {
	var sb strings.Builder
	sb.WriteString("MapAttr{")
	isFirstField := true
	for k, v := range a.data {
		if !isFirstField {
			sb.WriteString(", ")
		}

		fmt.Fprintf(&sb, "%#v", k)
		sb.WriteString(": ")
		switch a := v.(type) {
		case *StrMap:
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
