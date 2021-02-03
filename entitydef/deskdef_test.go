package entitydef

import (
	"encoding/json"
	"testing"

	. "github.com/go-playground/assert/v2"
	bson "go.mongodb.org/mongo-driver/bson"
)

func TestConvert(t *testing.T) {
	dd1 := EmptyDesk()
	dd2 := EmptyDesk()

	Equal(t, dd1.Equal(dd2), true)

	dd3 := NewDesk(100, 200, "300", 400, nil)
	dd4 := NewDesk(100, 200, "300", 400, nil)
	Equal(t, dd3.Equal(dd4), true)

	dd5 := NewDesk(100, 200, "300", 400, NewDesk(500, 600, "700", 800, nil))
	dd6 := NewDesk(100, 200, "300", 400, NewDesk(500, 600, "700", 800, nil))
	Equal(t, dd5.Equal(dd6), true)

}

func TestDesk(t *testing.T) {
	empty := EmptyDesk()
	Equal(t, empty.Data(), map[string]interface{}{
		"csvID":  int32(0),
		"height": int32(0),
		"name":   "",
		"width":  int32(0),
		"below":  (*Desk)(nil),
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

	empty.SetBelow(NewDesk(300, 200, "washington", 100, nil))
	Equal(t, empty.ChangeKey(), map[string]struct{}{"below": {}})
	Equal(t, empty.GetBelow().Data(), map[string]interface{}{
		"width":  int32(300),
		"height": int32(200),
		"csvID":  int32(100),
		"name":   "washington",
		"below":  (*Desk)(nil),
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
		newDesk := EmptyDesk()
		err = json.Unmarshal(bb, newDesk)
		Equal(t, err, nil)
		Equal(t, newDesk.Equal(empty), true)
	})

	t.Run("bson", func(t *testing.T) {
		empty := EmptyDesk()
		bb, err := bson.Marshal(empty)
		Equal(t, err, nil)
		newDesk := EmptyDesk()
		err = bson.Unmarshal(bb, newDesk)
		Equal(t, err, nil)
		Equal(t, newDesk.Equal(empty), true)
	})

}
