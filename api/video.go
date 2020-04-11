package api

import "github.com/gin-gonic/gin"

// CreateVideo 视频投稿
func CreateVideo(c *gin.Context) {
	service := service.CreateVideoService{}
	err := c.ShouldBind(service)
	if err != nil {
		c.JSON(200, ErrorResponse(err))
	} else {
		res := service.Create()
		c.JSON(200, res)
	}

}
