package entitydef

import (
	"encoding/json"
	"testing"

	. "github.com/go-playground/assert/v2"
	bson "go.mongodb.org/mongo-driver/bson"
)

func TestConvert(t *testing.T) {
	dd1 := EmptyDeskDef()
	dd2 := EmptyDeskDef()

	Equal(t, dd1.Equal(dd2), true)

	dd3 := NewDeskDef(100, 200, "300", 400, nil)
	dd4 := NewDeskDef(100, 200, "300", 400, nil)
	Equal(t, dd3.Equal(dd4), true)

	dd5 := NewDeskDef(100, 200, "300", 400, NewDeskDef(500, 600, "700", 800, nil))
	dd6 := NewDeskDef(100, 200, "300", 400, NewDeskDef(500, 600, "700", 800, nil))
	Equal(t, dd5.Equal(dd6), true)

}

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

	empty.SetBelow(NewDeskDef(300, 200, "washington", 100, nil))
	Equal(t, empty.ChangeKey(), map[string]struct{}{"below": {}})
	Equal(t, empty.GetBelow().data(), map[string]interface{}{
		"width":  int32(300),
		"height": int32(200),
		"csvID":  int32(100),
		"name":   "washington",
		"below":  (*DeskDef)(nil),
	})
	empty.ClearChangeKey()
	Equal(t, empty.HasChange(), false)
	Equal(t, empty.ChangeKey(), map[string]struct{}{})

	empty.GetBelow().SetHeight(10000)
	Equal(t, empty.HasChange(), true)
	Equal(t, empty.ChangeKey(), map[string]struct{}{"below": {}})

	t.Run("json", func(t *testing.T) {
		bb, err := json.Marshal(empty)
		Equal(t, err, nil)
		newDesk := EmptyDeskDef()
		err = json.Unmarshal(bb, newDesk)
		Equal(t, err, nil)
		Equal(t, newDesk.Equal(empty), true)
	})

	t.Run("bson", func(t *testing.T) {
		empty := EmptyDeskDef()
		bb, err := bson.Marshal(empty)
		Equal(t, err, nil)
		newDesk := EmptyDeskDef()
		err = bson.Unmarshal(bb, newDesk)
		Equal(t, err, nil)
		Equal(t, newDesk.Equal(empty), true)
	})

}
