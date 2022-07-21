package serializer

import "singo/model"

//type Scenery struct {
//	gorm.Model
//	Name          string
//	ViewType      string
//	Address       string
//	Price         int
//	ScenicAddress string
//	BusinessHour  string
//	ImageUrl      string
//	Detail        string
//	Longitude     float64
//	Latitude      float64
//	ScenicScore   string
//}

// History 用户序列化器
type Scenery struct {
	Name          string  `json:"name"`
	ViewType      string  `json:"view_type"`
	Address       string  `json:"address"`
	Price         int     `json:"price"`
	ScenicAddress string  `json:"scenic_address"`
	BusinessHour  string  `json:"business_hour"`
	ImageUrl      string  `json:"image_url"`
	Detail        string  `json:"detail"`
	Longitude     float64 `json:"longitude"`
	Latitude      float64 `json:"latitude"`
	ScenicScore   string  `json:"scenic_score"`
}

// BuildUser 序列化用户
func BuildScenery(item model.Scenery) Scenery {
	return Scenery{
		Name:          item.Name,
		ViewType:      item.ViewType,
		Address:       item.Address,
		Price:         item.Price,
		ScenicAddress: item.ScenicAddress,
		BusinessHour:  item.BusinessHour,
		ImageUrl:      item.ImageUrl,
		Detail:        item.Detail,
		Longitude:     item.Longitude,
		Latitude:      item.Latitude,
		ScenicScore:   item.ScenicScore,
	}
}

// BuildUserResponse 序列化用户响应
func BuildSceneryResponse(scenery model.Scenery) Response {
	return Response{
		Data: BuildScenery(scenery),
	}
}
