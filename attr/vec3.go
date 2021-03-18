package attr

import (
	"encoding/json"
	"sync"

	"go.mongodb.org/mongo-driver/bson"
)

type _Vec3 struct {
	X float32 `bson:"x" json:"x"`
	Y float32 `bson:"y" json:"y"`
	Z float32 `bson:"z" json:"z"`
}

func (a *_Vec3) Equal(other _Vec3) bool {
	return a.X == a.X && a.Y == a.Y && a.Z == a.Z
}

type Vec3 struct {
	parentKey string
	parent    Field

	data _Vec3
}

var vec3Pool *sync.Pool = &sync.Pool{
	New: func() interface{} {
		return &Vec3{
			// key:        "",
			parentKey: "",
			parent:    nil,
			data:      _Vec3{},
		}
	},
}

func EmptyVec3() *Vec3 {
	return NewVec3(0, 0, 0)
}
func NewVec3(x float32, y float32, z float32) *Vec3 {
	m := vec3Pool.Get().(*Vec3)
	m.data.X = x
	m.data.Y = y
	m.data.Z = z

	return m
}
func (a *Vec3) GetX() float32 {
	return a.data.X
}
func (a *Vec3) SetX(x float32) {
	a.data.X = x
	// if a.parent != nil {
	// 	a.parent.setChangeKey(a.parentKey)
	// }
}

func (a *Vec3) GetY() float32 {
	return a.data.Y
}
func (a *Vec3) SetY(y float32) {
	a.data.Y = y
	// if a.parent != nil {
	// 	a.parent.setChangeKey(a.parentKey)
	// }
}

func (a *Vec3) GetZ() float32 {
	return a.data.Z
}
func (a *Vec3) SetZ(z float32) {
	a.data.Z = z
	// if a.parent != nil {
	// 	a.parent.setChangeKey(a.parentKey)
	// }
}

func (a *Vec3) SetParent(k string, parent Field) {
	a.parentKey = k
	a.parent = parent
}

func (a *Vec3) Equal(other *Vec3) bool {
	return a.data.Equal(other.data)
}
func (a *Vec3) Undertype() interface{} {
	return a.data
}
func (a *Vec3) Data() map[string]interface{} {
	return map[string]interface{}{
		"X": a.data.X,
		"Y": a.data.Y,
		"Z": a.data.Z,
	}
}
func (a *Vec3) MarshalJSON() ([]byte, error) {
	return json.Marshal(a.data)
}
func (a *Vec3) UnmarshalJSON(b []byte) error {
	err := json.Unmarshal(b, &a.data)
	if err != nil {
		return err
	}
	return nil
}
func (a *Vec3) MarshalBSON() ([]byte, error) {
	return bson.Marshal(a.data)
}
func (a *Vec3) UnmarshalBSON(b []byte) error {
	err := bson.Unmarshal(b, &a.data)
	if err != nil {
		return err
	}
	return nil
}
