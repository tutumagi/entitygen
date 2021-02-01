package entitydef

import (
	"testing"

	. "github.com/go-playground/assert/v2"
)

func TestDeskDef(t *testing.T) {
	empty := EmptyDeskDef()
	Equal(t, empty.data(), map[string]interface{}{
		"csvID":  int32(0),
		"height": int32(0),
		"name":   "",
		"width":  int32(0),
		"below":  (*DeskDef)(nil),
	})
	Equal(t, empty.GetCsvID(), int32(0))
	Equal(t, empty.GetHeight(), int32(0))
	Equal(t, empty.GetWidth(), int32(0))
	Equal(t, empty.GetName(), "")
	Equal(t, empty.GetBelow(), nil)

	empty.SetCsvID(1001)
	Equal(t, empty.ChangeKey(), map[string]struct{}{"csvID": {}})
	empty.SetHeight(100)
	Equal(t, empty.ChangeKey(), map[string]struct{}{"csvID": {}, "height": {}})
	Equal(t, empty.HasChange(), true)

	empty.ClearChangeKey()
	Equal(t, empty.HasChange(), false)
	Equal(t, empty.ChangeKey(), map[string]struct{}{})

}
