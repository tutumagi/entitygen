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
		},
		Desks222: map[int32]*domain.Desk{
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

	desk := NewDeskDef(
		roomModel.Desk111.Width,
		roomModel.Desk111.Height,
		roomModel.Desk111.Name,
		roomModel.Desk111.CsvID,
	)

	deskss := EmptyKVInt32DeskDef()
	for k, v := range roomModel.Desks222 {
		deskss.Set(k, NewDeskDef(v.Width, v.Height, v.Name, v.CsvID))
	}

	room := NewRoomDef(
		roomModel.CsvPos,
		roomModel.BuildID,
		NewKVInt32Int32(roomModel.Extends),
		NewKVInt32Str(roomModel.Extends1),
		NewKVStrInt32(roomModel.Extends2),
		NewKVStrStr(roomModel.Extends3),
		desk,
		deskss,
	)

	return room, roomModel
}
