package model

import "gorm.io/gorm"

//旅游定制的结构体
type TravelCustomization struct {
	gorm.Model
	TraveCustomizationID int
	TravelMode           int    //多种方式对应int数据好处理
	Attraction           string //接string类型context输入
	DepartureTime        string //这里输入的是date
	ReturnTime           string //返程时间
}
