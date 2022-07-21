package api

import (
	"github.com/gin-gonic/gin"
	"singo/service"
)

//添加景点评论数据
func CreateUserTravel(c *gin.Context) {
	var service service.CreateUserTravelService
	if err := c.ShouldBind(&service); err == nil {
		res := service.Create()
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}
