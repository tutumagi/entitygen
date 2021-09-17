package attr

import (
	"fmt"
	"strings"
	"sync"
)

// StrMap str 作为 key 的 map
type StrMap struct {
	// key       interface{}
	parentKey string // 在父节点中的 key
	// 存这个数据的祖宗
	parent Field

	data map[string]interface{}

	// changedkey *sync.Map
	changedkey map[string]struct{}
}

func (a *StrMap) AOIData() *StrMap {
	return a
}

var strMapPool *sync.Pool = &sync.Pool{
	New: func() interface{} {
		return &StrMap{
			// key:        "",
			parentKey: "",
			parent:    nil,
			data:      map[string]interface{}{},
			// changedkey: &sync.Map{},
			changedkey: make(map[string]struct{}),
		}
	},
}

func NewStrMap(data map[string]interface{}) *StrMap {
	a := strMapPool.Get().(*StrMap)
	// a.key = ""
	a.parentKey = ""
	a.parent = nil

	// a.data["id"] = ""
	for k, v := range data {
		a.data[k] = v
	}

	return a
}

func ReleaseStrMap(strMap *StrMap) {
	for k := range strMap.data {
		delete(strMap.data, k)
	}
	// strMap.changedkey.Range(func(key, value interface{}) bool {
	// 	strMap.changedkey.Delete(key)
	// 	return true
	// })
	for k := range strMap.changedkey {
		delete(strMap.changedkey, k)
	}
	strMapPool.Put(strMap)
}

func (a *StrMap) SetParent(k string, parent Field) {
	if a == nil {
		return
	}
	if (a.parentKey != "" && a.parentKey != k) || (a.parent != nil && a.parent != parent) {
		panic(
			fmt.Sprintf(
				"has already set parent oldKey:%s newKey:%s oldParent:%v newParent:%v",
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
	if a == nil {
		return nil
	}
	result := map[string]interface{}{}
	for k, v := range a.data {
		if !isNil(v) {
			result[k] = v
		}
	}
	return result
	// return a.data
}

func (a *StrMap) FilterMap(filter func(k string) bool) map[string]interface{} {
	if a == nil {
		return nil
	}
	result := map[string]interface{}{}
	for k, v := range a.data {
		if filter(k) {
			result[k] = v
		}
	}
	return result
}

func (a *StrMap) ForEach(fn func(k string, v interface{}) bool) {
	for k, v := range a.data {
		// if k == "id" {
		// 	continue
		// }
		if !fn(k, v) {
			break
		}
	}
}

func (a *StrMap) Has(k string) bool {
	_, ok := a.data[k]
	return ok
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
	if a.data == nil {
		a.data = map[string]interface{}{}
	}
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

// Uint returns value with Uint type
func (a *StrMap) Uint(key string) uint {
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
func (a *StrMap) Uint8(key string) uint8 {
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
func (a *StrMap) Uint16(key string) uint16 {
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
func (a *StrMap) Uint32(key string) uint32 {
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
func (a *StrMap) Uint64(key string) uint64 {
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
		// NOTE: 注意，理论上来说这里不应该做判断，当没有 parent 的时候，必须要有 changedkey，如果没有代表有 bug
		// 但是这里 当我们使用 []*StrMap{} 去 mongo 里面做批量查询时，通过反射创建的StrMap里面 changedkey 为 nil
		if a.changedkey == nil {
			a.changedkey = make(map[string]struct{})
		}
		// a.changedkey.Store(key, struct{}{})
		a.changedkey[key] = struct{}{}
	} else {
		a.parent.setChangeKey(a.parentKey)
	}
}

func (a *StrMap) HasChange() bool {
	// hasChange := false
	// a.changedkey.Range(func(key, value interface{}) bool {
	// 	hasChange = true // 只要迭代了一次，就认为是有变化的 key
	// 	return false
	// })
	// return hasChange
	return len(a.changedkey) > 0
}

func (a *StrMap) ClearChangeKey() {
	for k := range a.changedkey {
		delete(a.changedkey, k)
	}
	// a.changedkey.Range(func(key, value interface{}) bool {
	// 	a.changedkey.Delete(key)
	// 	return true
	// })
}

func (a *StrMap) ChangeKey() map[string]struct{} {
	// var result = map[string]struct{}{}
	// a.changedkey.Range(func(key, value interface{}) bool {
	// 	result[key.(string)] = struct{}{}
	// 	return true
	// })
	// return result
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

// https://golang.org/ref/spec#Type_declarations
// type alias 和 type define 有不同
// type alias 语法 type MyInt = int.  MyInt 和 int 是完全一样的类型，共享一样的方法集，但是不能给 MyInt 添加自定义方法
// type define 语法 type MyInt int. MyInt 和 int 是完全独立的类型，MyInt 不共享 int 的方法集，可以给 MyInt 添加自定义方法
// 我们的生成工具用的 type define
func (a *StrMap) Equal(other *StrMap) bool {
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
			if im, ok := uu.Undertype().(_Vec3); ok {
				otherV := other.Value(k)
				if otherV != nil {
					if otherVV, ok := otherV.(IAttr); ok {
						if othervvv, ok := otherVV.Undertype().(_Vec3); ok {
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
func (a *StrMap) Data() map[string]interface{} {
	return a.data
}

func (a *StrMap) Undertype() interface{} {
	return a
}

func (a *StrMap) Len() int {
	return len(a.data)
}
