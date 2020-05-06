package service

import (
	"goVideo/model"
	"goVideo/serializer"
)

//ShowVideoDetail 视频详情
type ShowVideoDetail struct {
}

//show 展示视频
func (service *ShowVideoDetail) Show(id string) serializer.Response {
	video := model.Video{}
	// 根据主键查询一个记录
	err := model.DB.First(&video, id).Error
	if err == nil {

		video.AddView()

		return serializer.Response{
			Data: serializer.BuildVideo(video),
		}
	} else {
		return serializer.Response{
			Code:  404,
			Msg:   "视频不存在",
			Error: err.Error(),
		}
	}

}
