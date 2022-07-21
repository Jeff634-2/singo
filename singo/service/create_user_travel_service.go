package service

import (
	"singo/model"
	"singo/serializer"
)

//type UserTravel struct {
//	gorm.Model
//	//username做查询
//	UserID           int
//	UserName         string
//	UserSelectTravel []string //这个是接景点名字还是景点ID
//}

// CreateAttractionComment 景点添加评论服务
type CreateUserTravelService struct {
	UserID           int    `form:"userID" json:"userID" `
	UserName         string `form:"userName" json:"userName" `
	UserSelectTravel string `form:"userSelectTravel" json:"userSelectTravel" `
}

// CreateAttractionComment 创建旅游定制
func (service *CreateUserTravelService) Create() serializer.Response {

	//将数据存进数据库
	userTravel := model.UserTravel{
		UserID:           service.UserID,
		UserName:         service.UserName,
		UserSelectTravel: service.UserSelectTravel,
	}
	err := model.DB.Create(&userTravel).Error
	if err != nil {
		return serializer.Response{
			Code:  40001,
			Msg:   "添加评论失败",
			Error: err.Error(),
		}
	}
	return serializer.Response{
		Data: serializer.BuildUserTravel(userTravel),
	}
}
