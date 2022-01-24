package demodef

import (
	"github.com/g3n/engine/math32"
	"gitlab.testkaka.com/usm/game/entitygen/domain"
)

func mockRoom() (*RoomDef, *domain.RoomDef) {
	roomModel := &domain.RoomDef{
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
		Desk111: &domain.Desk{
			Width:  1024,
			Height: 768,
			Name:   "我是一张桌子",
			CsvID:  2048,
		},
		Desks222: map[int32]*domain.Desk{
			101: {
				Width:  101,
				Height: 1010,
				Name:   "desk one",
				CsvID:  10101,
			},
			102: {
				Width:  102,
				Height: 1020,
				Name:   "desk two",
				CsvID:  10201,
			},
		},
		Desks333: map[string]*domain.Desk{
			"mine": {
				Width:  2001,
				Height: 3001,
				Name:   "uber",
				CsvID:  4001,
			},
			"your": {
				Width:  5001,
				Height: 6001,
				Name:   "didi",
				CsvID:  7001,
			},
		},
		Int8ss: []int8{
			1, 2, 3, 4, 5, 6, 7, 8,
		},
		DesksArr: []*domain.Desk{
			{
				Width:  333,
				Height: 444,
				Name:   "innerdesk",
				CsvID:  333444,
			},
			{
				Width:  555,
				Height: 666,
				Name:   "wahaheihei",
				CsvID:  555666,
			},
		},
		Vec3: math32.NewVector3(1, 2, 3),
	}

	desk := NewDesk(
		roomModel.Desk111.Width,
		roomModel.Desk111.Height,
		roomModel.Desk111.Name,
		roomModel.Desk111.CsvID,
		EmptyDesk(),
	)

	int32desks := EmptyKVInt32Desk()
	for k, v := range roomModel.Desks222 {
		int32desks.Set(k, NewDesk(v.Width, v.Height, v.Name, v.CsvID, EmptyDesk()))
	}

	strdesks := EmptyKVStrDesk()
	for k, v := range roomModel.Desks333 {
		strdesks.Set(k, NewDesk(v.Width, v.Height, v.Name, v.CsvID, EmptyDesk()))
	}

	deskArr := EmptyDeskSlice()
	for _, v := range roomModel.DesksArr {
		deskArr.Add(NewDesk(v.Width, v.Height, v.Name, v.CsvID, EmptyDesk()))
	}

	room := NewRoomDef(
		roomModel.CsvPos,
		roomModel.BuildID,
		NewKVInt32Int32(roomModel.Extends),
		NewKVInt32Str(roomModel.Extends1),
		NewKVStrInt32(roomModel.Extends2),
		NewKVStrStr(roomModel.Extends3),
		desk,
		int32desks,
		strdesks,
		deskArr,
		NewInt8Slice(roomModel.Int8ss),
		NewVector3(roomModel.Vec3.X, roomModel.Vec3.Y, roomModel.Vec3.Z),
		NewVector2Slice([]*Vector2{NewVector2(1, 2), NewVector2(3, 4)}),
	)

	return room, roomModel
}
