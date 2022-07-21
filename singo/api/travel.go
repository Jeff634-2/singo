package api

import (
	"github.com/gin-gonic/gin"
	"singo/service"
)

//控制器做判断不做具体业务

// CreatetravelCustomization 添加旅游定制服务
func CreatetravelCustomization(c *gin.Context) {
	var service service.CreatetravelCustomizationService
	if err := c.ShouldBind(&service); err == nil {
		res := service.Create()
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}

//添加景点评论数据
func CreateAttractionComment(c *gin.Context) {
	var service service.CreateAttractionCommentService
	if err := c.ShouldBind(&service); err == nil {
		res := service.Create()
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}

//查看景点评论数据
func MyAttractionComment(c *gin.Context) {
	var service service.ShowMyAttractionCommentService
	res := service.Show(c.Param("id"))
	c.JSON(200, res)
}
