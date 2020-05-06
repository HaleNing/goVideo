package api

import (
	"github.com/gin-gonic/gin"
	"goVideo/service"
)

//DailyRank

func DailyRank(c *gin.Context) {
	rankService := service.DailyRankService{}
	//注意&
	if err := c.ShouldBind(&rankService); err == nil {
		res := rankService.Get()
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}

}
