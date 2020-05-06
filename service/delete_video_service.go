package service

import (
	"goVideo/model"
	"goVideo/serializer"
)

// DeleteVideo 删除视频服务

type DeleteVideo struct {
}

func (service *DeleteVideo) Delete(id string) serializer.Response {
	video := model.Video{}
	if len(id) == 0 {
		return serializer.Response{
			Code: 404,
			Msg:  "你提供了空信息",
		}
	}
	err := model.DB.First(&video, id).Error
	if err != nil {
		return serializer.Response{
			Code:  404,
			Msg:   "要删除的视频不存在",
			Error: err.Error(),
		}
	}
	if deleteErr := model.DB.Delete(&video).Error; deleteErr == nil {
		return serializer.Response{
			Data: serializer.BuildVideo(video),
		}
	} else {
		return serializer.Response{
			Code:  50003,
			Msg:   "视频删除失败",
			Error: deleteErr.Error(),
		}
	}

}
