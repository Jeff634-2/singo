package api

import (
	"github.com/gin-gonic/gin"
	"singo/service"
)

func MyTravelHistory(c *gin.Context) {
	var service service.ShowMyTravelHistoryService
	if err := c.ShouldBind(&service); err == nil {
		res := service.Show(c.Param("id"))
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}
