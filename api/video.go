package api

import (
	"github.com/gin-gonic/gin"
	"goVideo/service"
)

// CreateVideo 视频投稿接口
func CreateVideo(c *gin.Context) {
	user := CurrentUser(c)
	ser := service.CreateVideoService{}
	//注意&
	if err := c.ShouldBind(&ser); err == nil {
		res := ser.Create(user)
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
		limit := c.DefaultQuery("limit", 6)
		res := ser.ShowList(limit)
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}

}

// UpdateVideo 视频更新接口
func UpdateVideo(c *gin.Context) {
	ser := service.UpdateVideo{}
	//注意&
	if err := c.ShouldBind(&ser); err == nil {
		res := ser.Update(c.Param("id"))
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}

}

// DeleteVideo 视频删除接口
func DeleteVideo(c *gin.Context) {
	ser := service.DeleteVideo{}
	//注意&
	if err := c.ShouldBind(&ser); err == nil {
		res := ser.Delete(c.Param("id"))
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}

}

// TestLimit 测试limit接口
//func TestLimit(c *gin.Context)  {
//	ser := service.TestLimit{}
//	//注意&
//	if err := c.ShouldBind(&ser); err == nil {
//		limit := c.DefaultQuery("limit", "1")
//		res := ser.Limit(c.Param("id"),limit)
//		c.JSON(200, res)
//	} else {
//		c.JSON(200, ErrorResponse(err))
//	}
//}
