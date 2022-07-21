package model

import "gorm.io/gorm"

type UserTravel struct {
	gorm.Model
	//username做查询
	UserID           int
	UserName         string
	UserSelectTravel string //这个是接景点名字还是景点ID
}
