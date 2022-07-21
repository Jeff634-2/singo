package model

import "gorm.io/gorm"

type Attraction struct {
	gorm.Model
	//username做查询
	AttractionID int
	Title        string
	Address      string
	TypeId       int
	TagId        int
	Summary      string
	Content      string
	CreateTime   string
	ModifyTime   string
	Seen         string
	TravelerId   int
}
