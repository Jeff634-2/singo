package model

import "gorm.io/gorm"

// History 用户模型
type History struct {
	gorm.Model
	//username做查询
	UserName      string
	Departure     string
	Destination   string
	DepartureTime string
	ReturnTime    string
	TypeId        int
	TagId         int
	TotalCost     int
	Seen          string
	TravelerId    int
}
