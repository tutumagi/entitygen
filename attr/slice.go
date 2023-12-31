package attr

import (
	"fmt"
	"strings"
	"sync"
)

// Slice  slice
type Slice struct {
	data []interface{}

	// 存这个数据的父节点和在父节点中的 key
	parent    Field
	parentKey string
}

var strSlicePool *sync.Pool = &sync.Pool{
	New: func() interface{} {
		return &Slice{
			parentKey: "",
			data:      []interface{}{},

			parent: nil,
		}
	},
}

func NewSlice(data []interface{}) *Slice {
	strSlice := strSlicePool.Get().(*Slice)
	strSlice.parentKey = ""
	strSlice.parent = nil
	strSlice.data = data
	return strSlice
}

func ReleaseSlice(mm *Int32Map) {
	mm.data = map[int32]interface{}{}
	mm.parentKey = ""
	mm.parent = nil
	strSlicePool.Put(mm)
}

func (a *Slice) String() string {
	var sb strings.Builder
	sb.WriteString("Slice[")
	isFirstField := true
	for _, v := range a.data {
		if !isFirstField {
			sb.WriteString(", ")
		}

		fmt.Fprintf(&sb, "%v", v)

		isFirstField = false
	}
	sb.WriteString("]")
	return sb.String()
}

func (a *Slice) ToSlice() []interface{} {
	return a.data
}

func (a *Slice) ForEach(fn func(index int, v interface{}) bool) {
	for i, v := range a.data {
		if !fn(i, v) {
			break
		}
	}
}

func (a *Slice) change() {
	if a.parent != nil {
		a.parent.setChangeKey(a.parentKey)
	}
}

func (a *Slice) setChangeKey(k string) {
	a.change()
}

func (a *Slice) SetParent(k string, parent Field) {
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

func (a *Slice) SetData(data []interface{}) {
	a.data = data
}

func (a *Slice) Set(index int, val interface{}) {
	if index > len(a.data) {
		// TODO 这里是否需要打印出日志?
		return
	}
	a.data[index] = val
	// 这里缓存 修改的 key
	// 还有一种做法是 改变立马通知除去
	a.change()
}

func (a *Slice) Add(val interface{}) {
	a.data = append(a.data, val)
	a.change()
}

func (a *Slice) DeleteAt(index int) bool {
	if index < 0 || index >= len(a.data) {
		return false
	}
	src := a.data
	a.data = a.data[0:index]
	a.data = append(a.data, src[index+1:]...)

	a.change()
	return true
}

func (a *Slice) Len() int {
	return len(a.data)
}

func (a *Slice) Cap() int {
	return cap(a.data)
}

// Bool returns value with Bool type
func (a *Slice) Bool(index int) bool {
	if index >= len(a.data) {
		return false
	}

	r, ok := a.data[index].(bool)
	if !ok {
		return false
	}
	return r
}

// String returns value with String type
func (a *Slice) Str(index int) string {
	if index >= len(a.data) {
		return ""
	}
	r, ok := a.data[index].(string)
	if !ok {
		return ""
	}
	return r
}

// Value returns value with interface{} type
func (a *Slice) Value(index int) interface{} {
	if index >= len(a.data) {
		return nil
	}
	return a.data[index]
}

// Int returns value with Int type
func (a *Slice) Int(index int) int {
	if index >= len(a.data) {
		return 0
	}
	r, ok := a.data[index].(int)
	if !ok {
		return 0
	}
	return r
}

// Int8 returns value with Int8 type
func (a *Slice) Int8(index int) int8 {
	if index >= len(a.data) {
		return 0
	}
	r, ok := a.data[index].(int8)
	if !ok {
		return 0
	}
	return r
}

// Int16 returns value with Int16 type
func (a *Slice) Int16(index int) int16 {
	if index >= len(a.data) {
		return 0
	}
	r, ok := a.data[index].(int16)
	if !ok {
		return 0
	}
	return r
}

// Int32 returns value with Int32 type
func (a *Slice) Int32(index int) int32 {
	if index >= len(a.data) {
		return 0
	}
	r, ok := a.data[index].(int32)
	if !ok {
		return 0
	}
	return r
}

// Int64 returns value with Int64 type
func (a *Slice) Int64(index int) int64 {
	if index >= len(a.data) {
		return 0
	}
	r, ok := a.data[index].(int64)
	if !ok {
		return 0
	}
	return r
}

// UInt returns value with UInt type
func (a *Slice) UInt(index int) uint {
	if index >= len(a.data) {
		return 0
	}
	r, ok := a.data[index].(uint)
	if !ok {
		return 0
	}
	return r
}

// Uint8 returns value with Uint8 type
func (a *Slice) Uint8(index int) uint8 {
	if index >= len(a.data) {
		return 0
	}
	r, ok := a.data[index].(uint8)
	if !ok {
		return 0
	}
	return r
}

// Uint16 returns value with Uint16 type
func (a *Slice) Uint16(index int) uint16 {
	if index >= len(a.data) {
		return 0
	}
	r, ok := a.data[index].(uint16)
	if !ok {
		return 0
	}
	return r
}

// Uint32 returns value with Uint32 type
func (a *Slice) Uint32(index int) uint32 {
	if index >= len(a.data) {
		return 0
	}
	r, ok := a.data[index].(uint32)
	if !ok {
		return 0
	}
	return r
}

// Uint64 returns value with Uint64 type
func (a *Slice) Uint64(index int) uint64 {
	if index >= len(a.data) {
		return 0
	}
	r, ok := a.data[index].(uint64)
	if !ok {
		return 0
	}

	return r
}

// Float32 returns value with float32 type
func (a *Slice) Float32(index int) float32 {
	if index >= len(a.data) {
		return 0
	}
	r, ok := a.data[index].(float32)
	if !ok {
		return 0
	}
	return r
}

// Float64 returns value with float64 type
func (a *Slice) Float64(index int) float64 {
	if index >= len(a.data) {
		return 0
	}
	r, ok := a.data[index].(float64)
	if !ok {
		return 0
	}

	return r
}

func (a *Slice) Equal(other *Slice) bool {
	if a == nil && other == nil {
		return true
	}
	if (a == nil && other != nil) || (a != nil && other == nil) {
		return false
	}
	if a.Len() != other.Len() {
		return false
	}
	equal := true
	for i, v := range a.data {
		if uu, ok := v.(IAttr); ok {
			if im, ok := uu.Undertype().(*Int32Map); ok {
				otherV := other.Value(i)
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
				otherV := other.Value(i)
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
				otherV := other.Value(i)
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

		if v == other.data[i] {
			continue
		} else {
			equal = false
			break
		}
	}

	return equal
}

// 返回值只读，由外部自己保证不要去改这里面的东西
func (a *Slice) Data() []interface{} {
	return a.data
}

func (a *Slice) Undertype() interface{} {
	return a
}
