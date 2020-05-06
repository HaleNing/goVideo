package service

import (
	"fmt"
	"goVideo/cache"
	"goVideo/model"
	"goVideo/serializer"
	"strings"
)

type DailyRankService struct {
}

func (service *DailyRankService) Get() serializer.Response {
	var videos []model.Video
	results, _ := cache.RedisClient.ZRevRange(cache.DailyRankKey, 0, 9).Result()

	if len(results) > 1 {
		order := fmt.Sprintf("FIELD(id,%s)", strings.Join(results, ","))
		err := model.DB.Where("id in (?)", results).Order(order).Find(&videos).Error
		if err != nil {
			return serializer.Response{
				Code:  50000,
				Data:  "数据库连接错误",
				Error: err.Error(),
			}
		}
	}

	return serializer.Response{
		Data: serializer.BuildVideos(videos),
	}

}
