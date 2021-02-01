package entitydef

import (
	"encoding/json"
	"testing"

	"gitlab.gamesword.com/nut/entitygen/domain"

	. "github.com/go-playground/assert/v2"
	"go.mongodb.org/mongo-driver/bson"
)

func TestData(t *testing.T) {

	// 检查 数据
	t.Run("equal", func(t *testing.T) {
		room, roomModel := mockRoom()
		testEqualSource(t, room, roomModel)
	})

	// 检查 changekey
	t.Run("data-changekey", func(t *testing.T) {
		room, _ := mockRoom()
		testChangeKey(t, room)
	})

	t.Run("json-equal", func(t *testing.T) {
		room, roomModel := mockRoom()

		bb, err := json.Marshal(room)
		Equal(t, err, nil)
		newRoom := EmptyRoomDef()
		err = json.Unmarshal(bb, newRoom)
		Equal(t, err, nil)

		testEqualSource(t, newRoom, roomModel)
	})
}

func TestMarshalUnmarshal(t *testing.T) {
	t.Run("json", func(t *testing.T) {
		room, roomModel := mockRoom()
		testMarshalUnMarshal(t, json.Marshal, json.Unmarshal, room, roomModel)
	})

	t.Run("bson", func(t *testing.T) {
		room, roomModel := mockRoom()
		testMarshalUnMarshal(t, bson.Marshal, bson.Unmarshal, room, roomModel)
	})
}

func testEqualSource(t *testing.T, room *RoomDef, roomModel *domain.Room) {
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

	for k, v := range roomModel.Desks333 {
		d := room.GetDesks333().Get(k)
		Equal(t, d.GetHeight(), v.Height)
		Equal(t, d.GetWidth(), v.Width)
		Equal(t, d.GetName(), v.Name)
		Equal(t, d.GetCsvID(), v.CsvID)
	}

	for k, v := range roomModel.DesksArr {
		d := room.GetDesksArr().At(k)
		Equal(t, d.GetHeight(), v.Height)
		Equal(t, d.GetWidth(), v.Width)
		Equal(t, d.GetName(), v.Name)
		Equal(t, d.GetCsvID(), v.CsvID)
	}

	for k, v := range roomModel.Int8ss {
		d := room.GetInt8ss().At(k)
		Equal(t, d, v)
	}
}

func testEqualDef(t *testing.T, left *RoomDef, right *RoomDef) {
	Equal(t, left.GetCsvPos(), right.GetCsvPos())
	Equal(t, left.GetBuildID(), right.GetBuildID())
	Equal(t, left.GetExtends().Equal(right.GetExtends()), true)
	Equal(t, left.GetExtends1().Equal(right.GetExtends1()), true)
	Equal(t, left.GetExtends2().Equal(right.GetExtends2()), true)
	Equal(t, left.GetExtends3().Equal(right.GetExtends3()), true)
	Equal(t, left.GetDesk111().Equal(right.GetDesk111()), true)
	Equal(t, left.GetDesks222().Equal(right.GetDesks222()), true)
	Equal(t, left.GetDesks333().Equal(right.GetDesks333()), true)
	Equal(t, left.GetDesksArr().Equal(right.GetDesksArr()), true)
	Equal(t, left.GetInt8ss().Equal(right.GetInt8ss()), true)

	Equal(t, left.Equal(right), true)

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

	room.ClearChangeKey()
	Equal(t, room.ChangeKey(), map[string]struct{}{})
	room.GetDesksArr().At(0).SetHeight(3000)
	Equal(t, room.ChangeKey(), map[string]struct{}{"desks": {}})

	room.GetInt8ss().Set(3, 100)
	Equal(t, room.ChangeKey(), map[string]struct{}{"int8ss": {}, "desks": {}})

}

func testMarshalUnMarshal(
	t *testing.T,
	marshal func(v interface{}) ([]byte, error),
	unmarshal func([]byte, interface{}) error,
	room *RoomDef,
	model *domain.Room,
) {
	bbs, err := marshal(room)
	Equal(t, err, nil)

	newRoom := EmptyRoomDef()
	err = unmarshal(bbs, newRoom)
	Equal(t, err, nil)

	t.Run("equal", func(t *testing.T) {
		testEqualDef(t, room, newRoom)
	})

	t.Run("equal-old", func(t *testing.T) {
		testEqualSource(t, room, model)
	})

	t.Run("equal-new", func(t *testing.T) {
		testEqualSource(t, newRoom, model)
	})

	t.Run("oldroom", func(t *testing.T) {
		testChangeKey(t, room)
	})

	t.Run("newroom", func(t *testing.T) {
		testChangeKey(t, newRoom)
	})
}
