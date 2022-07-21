package serializer

import "singo/model"

type TravelCustomization struct {
	TraveCustomizationID int    `json:"id"`
	TravelMode           int    `json:"travelMode"`
	Attraction           string `json:"attraction"`
	DepartureTime        string `json:"departureTime"`
	ReturnTime           string `json:"returnTime"`
}

//type CreatetravelCustomizationService struct {
//	TravelMode    int    `form:"TravelMode" json:"TravelMode" binding:"required,min=1,max=30"`
//	Attraction    string `form:"Attraction" json:"Attraction" binding:"required,min=1,max=30"`
//	DepartureTime string `form:"departureTime" json:"departureTime" binding:"required,min=1,max=40"`
//	ReturnTime    string `form:"returnTime" json:"returnTime" binding:"required,min=1,max=40"`
//}

// BuildUser 序列化用户
func BuildTravelCustomization(item model.TravelCustomization) TravelCustomization {
	return TravelCustomization{
		TraveCustomizationID: item.TraveCustomizationID,
		TravelMode:           item.TravelMode,
		Attraction:           item.Attraction,
		DepartureTime:        item.DepartureTime,
		ReturnTime:           item.ReturnTime,
	}
}

//添加评论数据的序列化
type AttractionCommentDetail struct {
	AttractionName    string `json:"attractionName"`
	AttractionID      int    `json:"attractionID"`
	AttractionComment string `json:"attractionComment"`
}

//type CreatetravelCustomizationService struct {
//	TravelMode    int    `form:"TravelMode" json:"TravelMode" binding:"required,min=1,max=30"`
//	Attraction    string `form:"Attraction" json:"Attraction" binding:"required,min=1,max=30"`
//	DepartureTime string `form:"departureTime" json:"departureTime" binding:"required,min=1,max=40"`
//	ReturnTime    string `form:"returnTime" json:"returnTime" binding:"required,min=1,max=40"`
//}

// BuildUser 序列化用户
func BuildAttractionCommentDetail(item model.AttractionCommentDetail) AttractionCommentDetail {
	return AttractionCommentDetail{
		AttractionName:    item.AttractionName,
		AttractionID:      item.AttractionID,
		AttractionComment: item.AttractionComment,
	}
}
