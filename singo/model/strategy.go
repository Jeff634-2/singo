package model

import (
	"gorm.io/gorm"
)

type Strategy struct {
	gorm.Model
	AttractionName    string
	ImgUrl            string
	Price             int
	AttractionAddress string
	Detail            string
	Province          string
	Score             float64
	Like              int
	Type              string
}
