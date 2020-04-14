package service

import (
	"goVideo/model"
	"goVideo/serializer"
)

//ShowListVideo 查询视频列表
type ShowListVideo struct {
}

//ShowList 查询列表
func (service *ShowListVideo) ShowList() serializer.Response {
	var videos []model.Video
	// 查询所有记录
	err := model.DB.Find(&videos).Error
	if err == nil {
		return serializer.Response{
			Data: serializer.BuildVideos(videos),
		}
	} else {
		return serializer.Response{
			Code:  404,
			Msg:   "视频列表为空",
			Error: err.Error(),
		}
	}

}
