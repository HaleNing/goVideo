package service

import (
	"goVideo/model"
	"goVideo/serializer"
)

// TestLimit 测试limit接口
type TestLimit struct {
}

//  测试limit
func (service *TestLimit) Limit(id string, limit string) serializer.Response {
	var videos []model.Video
	// 查询所有记录
	err := model.DB.Limit(limit).Find(&videos).Error
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
