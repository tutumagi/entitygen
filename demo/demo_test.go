package demo

import (
	"encoding/json"
	"testing"

	. "github.com/go-playground/assert/v2"
	"go.mongodb.org/mongo-driver/bson"
)

func mockRoom() (*Room, *RoomXXX) {
	roomModel := &RoomXXX{
		// ID:      "1",
		CsvPos:  3,
		BuildID: "i am a build id",
		Extends: map[int32]int32{
			123: 456,
			789: 1011,
		},
		Extends1: map[int32]string{
			111: "hello",
			222: "world",
		},
		Extends2: map[string]int32{
			"tutu": 333,
			"fff":  444,
		},
		Extends3: map[string]string{
			"magi":   "jackie",
			"monica": "chen",
		},
		Desk: &DeskXXX{
			Width:  1024,
			Height: 768,
			Name:   "我是一张桌子",
		},
		Desks: map[int32]*DeskXXX{
			101: {
				Width:  101,
				Height: 1010,
				Name:   "desk one",
			},
			102: {
				Width:  102,
				Height: 1020,
				Name:   "desk two",
			},
		},
	}

	desk := NewDesk(
		roomModel.Desk.Width,
		roomModel.Desk.Height,
		roomModel.Desk.Name,
	)

	deskss := EmptyDesks()
	for k, v := range roomModel.Desks {
		deskss.Set(k, NewDesk(v.Width, v.Height, v.Name))
	}

	room := NewRoom(
		roomModel.CsvPos,
		roomModel.BuildID,
		NewRoomExtends(roomModel.Extends),
		NewRoomExtends1(roomModel.Extends1),
		NewRoomExtends2(roomModel.Extends2),
		NewRoomExtends3(roomModel.Extends3),
		desk,
		deskss,
	)

	return room, roomModel
}

func TestDemo(t *testing.T) {
	room, roomModel := mockRoom()
	// 检查 数据
	Equal(t, room.GetCsvPos(), roomModel.CsvPos)
	Equal(t, room.GetBuildID(), roomModel.BuildID)
	Equal(t, room.GetExtends().data(), roomModel.Extends)
	Equal(t, room.GetExtends1().data(), roomModel.Extends1)
	Equal(t, room.GetExtends2().data(), roomModel.Extends2)
	Equal(t, room.GetExtends3().data(), roomModel.Extends3)
	Equal(t, room.GetDesk().GetWidth(), roomModel.Desk.Width)
	Equal(t, room.GetDesk().GetHeight(), roomModel.Desk.Height)
	Equal(t, room.GetDesk().GetName(), roomModel.Desk.Name)

	for k, v := range roomModel.Desks {
		d := room.GetDesks().Get(k)
		Equal(t, d.GetHeight(), v.Height)
		Equal(t, d.GetWidth(), v.Width)
		Equal(t, d.GetName(), v.Name)
	}

	// 检查 changekey
	room.SetCsvPos(100)
	Equal(t, room._data.HasChange(), true)
	Equal(t, room._data.ChangeKey(), map[string]struct{}{"csv_pos": {}})
	room.SetBuildID("xxaabbcc")
	Equal(t, room._data.ChangeKey(), map[string]struct{}{
		"csv_pos":  {},
		"build_id": {},
	})
	room.SetExtends(NewRoomExtends(map[int32]int32{888: 999}))
	Equal(t, room._data.ChangeKey(), map[string]struct{}{
		"csv_pos":  {},
		"build_id": {},
		"extends":  {},
	})
	room.GetExtends1().Set(999, "money")
	Equal(t, room._data.ChangeKey(), map[string]struct{}{
		"csv_pos":  {},
		"build_id": {},
		"extends":  {},
		"extends1": {},
	})

	room._data.ClearChangeKey()
	Equal(t, room._data.HasChange(), false)
	Equal(t, room._data.ChangeKey(), map[string]struct{}{})

	room.GetExtends1().Set(999, "money")
	Equal(t, room._data.ChangeKey(), map[string]struct{}{
		"extends1": {},
	})

	room._data.ClearChangeKey()
	room.GetExtends1().Delete(999)
	Equal(t, room._data.ChangeKey(), map[string]struct{}{
		"extends1": {},
	})

	room._data.ClearChangeKey()
	// 这个 extends 没有这个 key，所以删掉后，没有这个 changkey
	room.GetExtends1().Delete(1000)
	Equal(t, room._data.ChangeKey(), map[string]struct{}{})

	room._data.ClearChangeKey()
	room.GetDesk().SetHeight(200)
	Equal(t, room.GetDesk().GetHeight(), int32(200))
	Equal(t, room._data.ChangeKey(), map[string]struct{}{"desk": {}})

	room._data.ClearChangeKey()
	Equal(t, room._data.ChangeKey(), map[string]struct{}{})
	room.GetDesks().Get(101).SetWidth(500)
	Equal(t, room._data.ChangeKey(), map[string]struct{}{"desks": {}})
}

func TestMarshalUnmarshal(t *testing.T) {
	room, _ := mockRoom()

	{
		bbs, err := json.Marshal(room)
		Equal(t, err, nil)

		newRoom := EmptyRoom()
		err = json.Unmarshal(bbs, newRoom)
		Equal(t, err, nil)

		Equal(t, newRoom.GetCsvPos(), room.GetCsvPos())
		Equal(t, newRoom.GetBuildID(), room.GetBuildID())
		Equal(t, newRoom.GetExtends().Equal(room.GetExtends()), true)
		Equal(t, newRoom.GetExtends1().Equal(room.GetExtends1()), true)
		Equal(t, newRoom.GetExtends2().Equal(room.GetExtends2()), true)
		Equal(t, newRoom.GetExtends3().Equal(room.GetExtends3()), true)
		Equal(t, newRoom.GetDesk().Equal(room.GetDesk()), true)
		Equal(t, newRoom.GetDesks().Equal(room.GetDesks()), true)

	}

	{
		bbs, err := bson.Marshal(room)
		Equal(t, err, nil)

		newRoom := EmptyRoom()
		err = bson.Unmarshal(bbs, newRoom)
		Equal(t, err, nil)

		Equal(t, newRoom.GetCsvPos(), room.GetCsvPos())
		Equal(t, newRoom.GetBuildID(), room.GetBuildID())
		Equal(t, newRoom.GetExtends().Equal(room.GetExtends()), true)
		Equal(t, newRoom.GetExtends1().Equal(room.GetExtends1()), true)
		Equal(t, newRoom.GetExtends2().Equal(room.GetExtends2()), true)
		Equal(t, newRoom.GetExtends3().Equal(room.GetExtends3()), true)
		Equal(t, newRoom.GetDesk().Equal(room.GetDesk()), true)
		Equal(t, newRoom.GetDesks().Equal(room.GetDesks()), true)
	}
}
