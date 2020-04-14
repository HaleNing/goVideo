package api

import (
	"github.com/gin-gonic/gin"
	"goVideo/service"
)

// CreateVideo 视频投稿接口
func CreateVideo(c *gin.Context) {
	ser := service.CreateVideoService{}
	//注意&
	if err := c.ShouldBind(&ser); err == nil {
		res := ser.Create()
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}

}

// ShowVideo 视频详情接口
func ShowVideo(c *gin.Context) {
	ser := service.ShowVideoDetail{}
	//注意&
	if err := c.ShouldBind(&ser); err == nil {
		res := ser.Show(c.Param("id"))
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}

}

// ListVideo 视频列表接口
func ListVideo(c *gin.Context) {
	ser := service.ShowListVideo{}
	//注意&
	if err := c.ShouldBind(&ser); err == nil {
		res := ser.ShowList()
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}

}
