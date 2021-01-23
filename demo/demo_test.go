package demo

import (
	"encoding/json"
	"testing"

	. "github.com/go-playground/assert/v2"
	"go.mongodb.org/mongo-driver/bson"
)

func TestDemo(t *testing.T) {
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
	}

	room := NewRoom(
		roomModel.CsvPos,
		roomModel.BuildID,
		roomModel.Extends,
		roomModel.Extends1,
	)

	{
		bbs, err := json.Marshal(room)
		Equal(t, err, nil)

		newRoom := NewRoom(0, "", nil, nil)
		err = json.Unmarshal(bbs, newRoom)
		Equal(t, err, nil)

		Equal(t, newRoom.GetCsvPos(), room.GetCsvPos())
		Equal(t, newRoom.GetBuildID(), room.GetBuildID())
		Equal(t, newRoom.GetExtends().Equal(room.GetExtends()), true)
		Equal(t, newRoom.GetExtends1().Equal(room.GetExtends1()), true)
	}

	{
		bbs, err := bson.Marshal(room)
		Equal(t, err, nil)

		newRoom := NewRoom(0, "", nil, nil)
		err = bson.Unmarshal(bbs, newRoom)
		Equal(t, err, nil)

		Equal(t, newRoom.GetCsvPos(), room.GetCsvPos())
		Equal(t, newRoom.GetBuildID(), room.GetBuildID())
		Equal(t, newRoom.GetExtends().Equal(room.GetExtends()), true)
		Equal(t, newRoom.GetExtends1().Equal(room.GetExtends1()), true)
	}

}
