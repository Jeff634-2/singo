package api

import (
	"github.com/gin-gonic/gin"
	"singo/service"
)

//查看景点评论数据
func UserRsView(c *gin.Context) {
	//处理推荐的一些逻辑

	var service service.ShowMyRsView
	res := service.Show(c.Param("id"))
	c.JSON(200, res)
}
