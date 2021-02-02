package entitydef

import (
	"testing"

	. "github.com/go-playground/assert/v2"
)

func Test_EmptyDeskDefSlice(t *testing.T) {
	empty := EmptyDeskDefSlice()

	Equal(t, empty.Count(), 0)

	dd1 := NewDeskDef(100, 101, "102", 103, nil)
	empty.Add(dd1)
	Equal(t, empty.Count(), 1)
	Equal(t, empty.At(0).Equal(dd1), true)
	Equal(t, empty.At(1), nil)

	dd2 := NewDeskDef(200, 201, "202", 203, nil)
	empty.Add(dd2)
	Equal(t, empty.Count(), 2)
	Equal(t, empty.At(0).Equal(dd1), true)
	Equal(t, empty.At(1).Equal(dd2), true)
	Equal(t, empty.At(2), nil)

	ddReplace2 := NewDeskDef(300, 302, "303", 304, nil)
	empty.Set(1, ddReplace2)
	Equal(t, empty.Count(), 2)
	Equal(t, empty.At(0).Equal(dd1), true)
	Equal(t, empty.At(1).Equal(dd2), false)
	Equal(t, empty.At(1).Equal(ddReplace2), true)
	Equal(t, empty.At(2), nil)

	empty.DelAt(0)
	Equal(t, empty.Count(), 1)
	Equal(t, empty.At(0).Equal(dd2), false)
	Equal(t, empty.At(0).Equal(ddReplace2), true)
	Equal(t, empty.At(1), nil)

	empty.DelAt(0)
	Equal(t, empty.Count(), 0)
	Equal(t, empty.At(0), nil)

	empty.DelAt(0)
	Equal(t, empty.Count(), 0)
	Equal(t, empty.At(0), nil)
}

func Test_DeskDefSlice(t *testing.T) {
	item0 := NewDeskDef(100, 101, "first", 102, NewDeskDef(202, 201, "first-first", 202, nil))
	empty := NewDeskDefSlice([]*DeskDef{
		item0,
	})

	Equal(t, empty.Count(), 1)
	Equal(t, empty.At(0).Equal(item0), true)

	dd1 := NewDeskDef(100, 101, "102", 103, nil)
	empty.Add(dd1)
	Equal(t, empty.Count(), 2)
	Equal(t, empty.At(0).Equal(item0), true)
	Equal(t, empty.At(1).Equal(dd1), true)
	Equal(t, empty.At(2), nil)

	dd2 := NewDeskDef(200, 201, "202", 203, nil)
	empty.Add(dd2)
	Equal(t, empty.Count(), 3)
	Equal(t, empty.At(0).Equal(item0), true)
	Equal(t, empty.At(1).Equal(dd1), true)
	Equal(t, empty.At(2).Equal(dd2), true)
	Equal(t, empty.At(3), nil)

	ddReplace1 := NewDeskDef(300, 302, "303", 304, nil)
	empty.Set(1, ddReplace1)
	Equal(t, empty.Count(), 3)
	Equal(t, empty.At(0).Equal(item0), true)
	Equal(t, empty.At(1).Equal(dd1), false)
	Equal(t, empty.At(1).Equal(ddReplace1), true)
	Equal(t, empty.At(2).Equal(dd2), true)
	Equal(t, empty.At(3), nil)

	empty.DelAt(0)
	Equal(t, empty.Count(), 2)
	Equal(t, empty.At(0).Equal(item0), false)
	Equal(t, empty.At(0).Equal(ddReplace1), true)
	Equal(t, empty.At(1).Equal(dd2), true)
	Equal(t, empty.At(2), nil)

	empty.DelAt(0)
	Equal(t, empty.Count(), 1)
	Equal(t, empty.At(0).Equal(dd2), true)
	Equal(t, empty.At(1), nil)

	empty.DelAt(0)
	Equal(t, empty.Count(), 0)
	Equal(t, empty.At(0), nil)
}
