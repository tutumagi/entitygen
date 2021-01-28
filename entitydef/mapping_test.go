package entitydef

import (
	"encoding/json"
	"testing"

	. "github.com/go-playground/assert/v2"
	"go.mongodb.org/mongo-driver/bson"
)

func TestData(t *testing.T) {
	room, roomModel := mockRoom()
	// 检查 数据
	Equal(t, room.GetCsvPos(), roomModel.CsvPos)
	Equal(t, room.GetBuildID(), roomModel.BuildID)
	Equal(t, room.GetExtends().data(), roomModel.Extends)
	Equal(t, room.GetExtends1().data(), roomModel.Extends1)
	Equal(t, room.GetExtends2().data(), roomModel.Extends2)
	Equal(t, room.GetExtends3().data(), roomModel.Extends3)
	Equal(t, room.GetDesk111().GetWidth(), roomModel.Desk111.Width)
	Equal(t, room.GetDesk111().GetHeight(), roomModel.Desk111.Height)
	Equal(t, room.GetDesk111().GetName(), roomModel.Desk111.Name)
	Equal(t, room.GetDesk111().GetCsvID(), roomModel.Desk111.CsvID)

	for k, v := range roomModel.Desks222 {
		d := room.GetDesks222().Get(k)
		Equal(t, d.GetHeight(), v.Height)
		Equal(t, d.GetWidth(), v.Width)
		Equal(t, d.GetName(), v.Name)
		Equal(t, d.GetCsvID(), v.CsvID)
	}

	// 检查 changekey
	t.Run("data-changekey", func(t *testing.T) {
		testChangeKey(t, room)
	})
}

func testChangeKey(t *testing.T, room *RoomDef) {
	room.ClearChangeKey()
	Equal(t, room.HasChange(), false)
	Equal(t, room.ChangeKey(), map[string]struct{}{})

	room.SetCsvPos(100)
	Equal(t, room.HasChange(), true)
	Equal(t, room.ChangeKey(), map[string]struct{}{"csv_pos": {}})
	room.SetBuildID("xxaabbcc")
	Equal(t, room.GetBuildID(), "xxaabbcc")
	Equal(t, room.ChangeKey(), map[string]struct{}{
		"csv_pos":  {},
		"build_id": {},
	})
	room.SetExtends(NewKVInt32Int32(map[int32]int32{888: 999}))
	Equal(t, room.GetExtends().data(), map[int32]int32{888: 999})
	Equal(t, room.ChangeKey(), map[string]struct{}{
		"csv_pos":  {},
		"build_id": {},
		"extends":  {},
	})
	room.GetExtends1().Set(999, "money")
	Equal(t, room.GetExtends1().Get(999), "money")
	Equal(t, room.ChangeKey(), map[string]struct{}{
		"csv_pos":  {},
		"build_id": {},
		"extends":  {},
		"extends1": {},
	})

	room.ClearChangeKey()
	Equal(t, room.HasChange(), false)
	Equal(t, room.ChangeKey(), map[string]struct{}{})

	room.GetExtends1().Set(999, "moneykkk")
	Equal(t, room.GetExtends1().Get(999), "moneykkk")
	Equal(t, room.ChangeKey(), map[string]struct{}{
		"extends1": {},
	})

	room.ClearChangeKey()
	room.GetExtends1().Delete(999)
	Equal(t, room.GetExtends1().Get(999), "")
	Equal(t, room.ChangeKey(), map[string]struct{}{
		"extends1": {},
	})

	room.ClearChangeKey()
	// 这个 extends 没有这个 key，所以删掉后，没有这个 changkey
	room.GetExtends1().Delete(1000)
	Equal(t, room.ChangeKey(), map[string]struct{}{})

	room.ClearChangeKey()
	room.GetDesk111().SetHeight(200)
	Equal(t, room.GetDesk111().GetHeight(), int32(200))
	Equal(t, room.ChangeKey(), map[string]struct{}{"desk888": {}})

	room.ClearChangeKey()
	Equal(t, room.ChangeKey(), map[string]struct{}{})
	room.GetDesks222().Get(101).SetWidth(500)
	Equal(t, room.ChangeKey(), map[string]struct{}{"desks999": {}})
}

func TestMarshalUnmarshal(t *testing.T) {
	t.Run("json", func(t *testing.T) {
		room, _ := mockRoom()

		bbs, err := json.Marshal(room)
		Equal(t, err, nil)

		newRoom := EmptyRoomDef()
		err = json.Unmarshal(bbs, newRoom)
		Equal(t, err, nil)

		Equal(t, newRoom.GetCsvPos(), room.GetCsvPos())
		Equal(t, newRoom.GetBuildID(), room.GetBuildID())
		Equal(t, newRoom.GetExtends().Equal(room.GetExtends()), true)
		Equal(t, newRoom.GetExtends1().Equal(room.GetExtends1()), true)
		Equal(t, newRoom.GetExtends2().Equal(room.GetExtends2()), true)
		Equal(t, newRoom.GetExtends3().Equal(room.GetExtends3()), true)
		Equal(t, newRoom.GetDesk111().Equal(room.GetDesk111()), true)
		Equal(t, newRoom.GetDesks222().Equal(room.GetDesks222()), true)

		Equal(t, room.ChangeKey(), map[string]struct{}{})
		Equal(t, newRoom.ChangeKey(), map[string]struct{}{})

		t.Run("oldroom", func(t *testing.T) {
			testChangeKey(t, room)
		})

		t.Run("newroom", func(t *testing.T) {
			testChangeKey(t, newRoom)
		})
	})

	t.Run("bson", func(t *testing.T) {
		room, _ := mockRoom()

		bbs, err := bson.Marshal(room)
		Equal(t, err, nil)

		newRoom := EmptyRoomDef()
		err = bson.Unmarshal(bbs, newRoom)
		Equal(t, err, nil)

		Equal(t, newRoom.GetCsvPos(), room.GetCsvPos())
		Equal(t, newRoom.GetBuildID(), room.GetBuildID())
		Equal(t, newRoom.GetExtends().Equal(room.GetExtends()), true)
		Equal(t, newRoom.GetExtends1().Equal(room.GetExtends1()), true)
		Equal(t, newRoom.GetExtends2().Equal(room.GetExtends2()), true)
		Equal(t, newRoom.GetExtends3().Equal(room.GetExtends3()), true)
		Equal(t, newRoom.GetDesk111().Equal(room.GetDesk111()), true)
		Equal(t, newRoom.GetDesks222().Equal(room.GetDesks222()), true)

		t.Run("oldroom", func(t *testing.T) {
			testChangeKey(t, room)
		})

		t.Run("newroom", func(t *testing.T) {
			testChangeKey(t, newRoom)
		})
	})
}
