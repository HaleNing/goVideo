package service

import (
	"goVideo/model"
	"goVideo/serializer"
)

//UpdateVideo 视频更新服务

/**
选择后续更新不可以更改视频标题
*/
type UpdateVideo struct {
	Info string `form:"info" json:"info" binding:"required,min=0,max=200"`
}

//Update 更新
func (service *UpdateVideo) Update(id string) serializer.Response {
	video := model.Video{}
	err := model.DB.First(&video, id).Error
	if err != nil {
		return serializer.Response{
			Code:  404,
			Msg:   "视频不存在",
			Error: err.Error(),
		}
	}
	updateErr := model.DB.Model(&video).Update("info", service.Info).Error
	if updateErr == nil {
		return serializer.Response{
			Data: serializer.BuildVideo(video),
		}
	} else {
		return serializer.Response{
			Code:  50000,
			Msg:   "视频更新失败",
			Error: updateErr.Error(),
		}
	}
}
