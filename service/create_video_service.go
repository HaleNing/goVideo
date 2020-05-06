package service

import (
	"goVideo/model"
	"goVideo/serializer"
)

//CreateVideoService 视频投稿服务
type CreateVideoService struct {
	TitleName string `formorm:"title_name" json:"title_name" binding:"required,min=5,max=30"`
	Info      string `form:"info" json:"info" binding:"required,min=0,max=200"`
	Url       string `form:"url" json:"url" binding:"required,min=0,max=200"`
}

//Create 创建视频
func (service *CreateVideoService) Create(user *model.User) serializer.Response {
	video := model.Video{
		TitleName: service.TitleName,
		Info:      service.Info,
		Url:       service.Url,
		UserId:    user.ID,
	}
	//在数据库中创建一条视频 注意&
	err := model.DB.Create(&video).Error
	if err == nil {
		return serializer.Response{
			Data: serializer.BuildVideo(video),
		}
	} else {
		return serializer.Response{
			Code:  50001,
			Msg:   "视频创建失败",
			Error: err.Error(),
		}
	}
}
