package entitydef

import "entitygen/domain"

func mockRoom() (*RoomDef, *domain.Room) {
	roomModel := &domain.Room{
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
			"mine": &domain.Desk{
				Width:  2001,
				Height: 3001,
				Name:   "uber",
				CsvID:  4001,
			},
			"your": &domain.Desk{
				Width:  5001,
				Height: 6001,
				Name:   "didi",
				CsvID:  7001,
			},
		},
	}

	desk := NewDeskDef(
		roomModel.Desk111.Width,
		roomModel.Desk111.Height,
		roomModel.Desk111.Name,
		roomModel.Desk111.CsvID,
	)

	int32desks := EmptyKVInt32DeskDef()
	for k, v := range roomModel.Desks222 {
		int32desks.Set(k, NewDeskDef(v.Width, v.Height, v.Name, v.CsvID))
	}

	strdesks := EmptyKVStrDeskDef()
	for k, v := range roomModel.Desks333 {
		strdesks.Set(k, NewDeskDef(v.Width, v.Height, v.Name, v.CsvID))
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
	)

	return room, roomModel
}
