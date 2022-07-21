package serializer

import (
	"singo/model"
)

//type UserTravel struct {
//	gorm.Model
//	//username做查询
//	UserID           int
//	UserName         string
//	UserSelectTravel []string //这个是接景点名字还是景点ID
//}

type UserTravel struct {
	UserID           int    `json:"userID"`
	UserName         string `json:"userName"`
	UserSelectTravel string `json:"userSelectTravel"`
}

//type CreatetravelCustomizationService struct {
//	TravelMode    int    `form:"TravelMode" json:"TravelMode" binding:"required,min=1,max=30"`
//	Attraction    string `form:"Attraction" json:"Attraction" binding:"required,min=1,max=30"`
//	DepartureTime string `form:"departureTime" json:"departureTime" binding:"required,min=1,max=40"`
//	ReturnTime    string `form:"returnTime" json:"returnTime" binding:"required,min=1,max=40"`
//}

// BuildUser 序列化用户
func BuildUserTravel(item model.UserTravel) UserTravel {
	return UserTravel{
		UserID:           item.UserID,
		UserName:         item.UserName,
		UserSelectTravel: item.UserSelectTravel,
	}
}
