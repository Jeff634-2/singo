package serializer

import (
	"singo/model"
)

// History 用户序列化器
type History struct {
	UserName      string `json:"user_name"`
	Departure     string `json:"departure"`
	Destination   string `json:"destination"`
	DepartureTime string `json:"departureTime"`
	ReturnTime    string `json:"returnTime"`
	TypeId        int    `json:"typeid"`
	TagId         int    `json:"tagid"`
	TotalCost     int    `json:"totalCost"`
	Seen          string `json:"seen"`
	TravelerId    int    `json:"travelerId"`
}

// BuildUser 序列化用户
func BuildUserHistory(item model.History) History {
	return History{
		UserName:      item.UserName,
		Departure:     item.Departure,
		Destination:   item.Destination,
		DepartureTime: item.DepartureTime,
		ReturnTime:    item.ReturnTime,
		TypeId:        item.TypeId,
		TagId:         item.TagId,
		TotalCost:     item.TotalCost,
		Seen:          item.Seen,
		TravelerId:    item.TravelerId,
	}
}

// BuildUserResponse 序列化用户响应
func BuildUserHistoryResponse(history model.History) Response {
	return Response{
		Data: BuildUserHistory(history),
	}
}
