package model

type Like struct {
	ObjectMeta
	UserId string `json:"userid" gorm:"column:userid"`
	ItemId string `json:"itemid" gorm:"column:userid"`
}

type LikeList struct {
	ListMeta
	Items []Like
}
