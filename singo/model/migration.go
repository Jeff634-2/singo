package model

//执行数据迁移

func migration() {
	// 自动迁移模式
	_ = DB.AutoMigrate(&User{})
	_ = DB.AutoMigrate(&TravelCustomization{})
	_ = DB.AutoMigrate(&History{})
	_ = DB.AutoMigrate(&Strategy{})
	_ = DB.AutoMigrate(&AttractionCommentDetail{})
	_ = DB.AutoMigrate(&UserTravel{})
}
