package service

import (
	"github.com/jinzhu/gorm"
	"goVideo/model"
	"goVideo/serializer"
)

//CreateVideoService 视频投稿服务
type CreateVideoService struct {
	TitleName string `form:"title_name" json:"title_name" binding:"required,min=5,max=30"`
	Info      string `form:"info" json:"info" binding:"required,min=0,max=200"`
}

//Create 创建视频
func (service *CreateVideoService) Create() serializer.Response {
	video := model.Video{
		Model:     gorm.Model{},
		TitleName: service.TitleName,
		Info:      service.Info,
	}
	//model.DB.Create(video)
}
