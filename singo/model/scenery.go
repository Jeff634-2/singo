package model

import "gorm.io/gorm"

type Scenery struct {
	gorm.Model
	Name          string
	ViewType      string
	Address       string
	Price         int
	ScenicAddress string
	BusinessHour  string
	ImageUrl      string
	Detail        string
	Longitude     float64
	Latitude      float64
	ScenicScore   string
	viewType      string
}
