package domain

type LandAttr struct {
	ID          string
	LandIdx     int32
	Price       int32
	OwnEntities map[string]string
}
