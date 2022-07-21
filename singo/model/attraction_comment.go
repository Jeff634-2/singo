package model

import "gorm.io/gorm"

type AttractionCommentDetail struct {
	gorm.Model
	AttractionName    string
	AttractionID      int
	AttractionComment string
}
