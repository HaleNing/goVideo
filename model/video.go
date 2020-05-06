package model

import (
	"github.com/jinzhu/gorm"
	"goVideo/cache"
	"strconv"
)

// Video 视频模型
type Video struct {
	gorm.Model
	TitleName string
	Info      string
	Url       string
	UserId    uint
}

// 返回视频点击数
func (video *Video) View() uint64 {
	result, _ := cache.RedisClient.Get(cache.VideoViewKey(video.ID)).Result()
	parseUint, _ := strconv.ParseUint(result, 10, 64)
	return parseUint
}

// 对视频的操作
func (video *Video) AddView() {
	// 增加视频点击数
	cache.RedisClient.Incr(cache.VideoViewKey(video.ID))
	// 增加排行点击数
	cache.RedisClient.ZIncrBy(cache.DailyRankKey, 1, strconv.Itoa(int(video.ID)))
}
