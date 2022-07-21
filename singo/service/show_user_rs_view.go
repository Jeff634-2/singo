package service

import (
	"singo/model"
	"singo/serializer"
)

type ShowMyRsView struct {
}

/*
完成的任务是
1.找出用户对应的view_type对应的景点列表
2.以shuffle等策略输出景点ID的形式输出
*/
func (service *ShowMyRsView) Show(viewType string) serializer.Response {
	//找出对应id为1的user喜欢的

	viewType = "主题乐园"

	var myScenery model.Scenery
	err := model.DB.First(&myScenery, viewType).Error
	//err := model.DB.Where(&model.Scenery{viewType: viewType}).Find(&myScenery)
	//db.Where(&User{Name: "jinzhu", Age: 0}).Find(&users)

	if err != nil {
		return serializer.Response{
			Code:  40001,
			Msg:   "展示用户推荐失败",
			Error: "404",
		}
	}

	return serializer.Response{
		Data: serializer.BuildScenery(myScenery),
	}
}
