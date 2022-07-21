package service

import (
	"singo/model"
	"singo/serializer"
)

// CreatetravelCustomizationService 旅游定制服务
type CreatetravelCustomizationService struct {
	TraveCustomizationID int    `form:"traveCustomizationID" json:"traveCustomizationID" `
	TravelMode           int    `form:"travelMode" json:"travelMode" `
	Attraction           string `form:"attraction" json:"attraction" `
	DepartureTime        string `form:"departureTime" json:"departureTime" `
	ReturnTime           string `form:"returnTime" json:"returnTime" `
}

// Create 创建旅游定制
func (service *CreatetravelCustomizationService) Create() serializer.Response {

	//将数据存进数据库
	travelCustomization := model.TravelCustomization{
		TraveCustomizationID: service.TraveCustomizationID,
		TravelMode:           service.TravelMode,
		Attraction:           service.Attraction,
		DepartureTime:        service.DepartureTime,
		ReturnTime:           service.ReturnTime,
	}
	err := model.DB.Create(&travelCustomization).Error
	if err != nil {
		return serializer.Response{
			Code:  40001,
			Msg:   "旅游定制失败",
			Error: err.Error(),
		}
	}
	return serializer.Response{
		Data: serializer.BuildTravelCustomization(travelCustomization),
	}
}

// CreateAttractionComment 景点添加评论服务
type CreateAttractionCommentService struct {
	AttractionName    string `form:"attractionName" json:"attractionName" `
	AttractionID      int    `form:"attractionID" json:"attractionID" `
	AttractionComment string `form:"attractionComment" json:"attractionComment" `
}

// CreateAttractionComment 创建旅游定制
func (service *CreateAttractionCommentService) Create() serializer.Response {

	//将数据存进数据库
	attractionCommentDetail := model.AttractionCommentDetail{
		AttractionName:    service.AttractionName,
		AttractionID:      service.AttractionID,
		AttractionComment: service.AttractionComment,
	}
	err := model.DB.Create(&attractionCommentDetail).Error
	if err != nil {
		return serializer.Response{
			Code:  40001,
			Msg:   "添加评论失败",
			Error: err.Error(),
		}
	}
	return serializer.Response{
		Data: serializer.BuildAttractionCommentDetail(attractionCommentDetail),
	}
}

//查看评论数据
type ShowMyAttractionCommentService struct {
}

// Show 展示用户历史服务
func (service *ShowMyAttractionCommentService) Show(id string) serializer.Response {

	var MyAttractionComment model.AttractionCommentDetail
	err := model.DB.First(&MyAttractionComment, id).Limit(10).Error

	if err != nil {
		return serializer.Response{
			Code:  40001,
			Msg:   "展示景点旅游评论数据失败",
			Error: err.Error(),
		}
	}

	return serializer.Response{
		Data: serializer.BuildAttractionCommentDetail(MyAttractionComment),
	}
}
