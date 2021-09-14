package entity

type DataItem struct {
	Key string      `bson:"key" json:"key"`
	Val interface{} `bson:"val" json:"val"`
}

type DataField struct {
	Name string `bson:"name" json:"name"`
	Data string `bson:"data" json:"data"`
}
