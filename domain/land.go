package domain

type LandAttr struct {
	ID      string `bson:"id" json:"id"`
	LandIdx int32  `bson:"land_idx" json:"land_idx"`
	Price   int32  `bson:"price" json:"price"`
	// OwnEntities *map[string]string
}
