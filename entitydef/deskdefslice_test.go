package entitydef

import (
	"testing"

	. "github.com/go-playground/assert/v2"
)

func TestDeskDefSlice(t *testing.T) {
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
}
