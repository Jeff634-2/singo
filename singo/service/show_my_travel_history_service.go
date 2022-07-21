package service

import (
	"singo/model"
	"singo/serializer"
)

type ShowMyTravelHistoryService struct {
}

// Show 展示用户历史服务
func (service *ShowMyTravelHistoryService) Show(id string) serializer.Response {

	var myTravelHistory model.History
	err := model.DB.First(&myTravelHistory, id).Error

	if err != nil {
		return serializer.Response{
			Code:  40001,
			Msg:   "展示用户旅游历史失败",
			Error: err.Error(),
		}
	}

	return serializer.Response{
		Data: serializer.BuildUserHistory(myTravelHistory),
	}
}
