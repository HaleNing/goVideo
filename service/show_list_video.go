package service

import (
	"goVideo/model"
	"goVideo/serializer"
)

//ShowListVideo 查询视频列表
type ShowListVideo struct {
	Limit int `form:"limit"`
	start int `form:"start"`
}

//ShowList 查询列表
func (service *ShowListVideo) ShowList(limit int) serializer.Response {
	var videos []model.Videos
	total := 0
	if service.Limit == 0 {
		service.Limit = limit
	}
	if err := model.DB.Model(model.Video{}).Count(&total).Error; err != nil {
		return serializer.Response{
			Code:  50000,
			Msg:   "数据库连接错误",
			Error: err.Error(),
		}
	}

	// 查询所有记录
	if err := model.DB.Limit(service.Limit).Offset(service.start).Find(&videos).Error; err != nil {
		return serializer.Response{
			Code:  50000,
			Msg:   "数据库连接错误",
			Error: err.Error(),
		}
	}
	return serializer.BuildListResponse(serializer.BuildVideos(videos), uint(total))
}
