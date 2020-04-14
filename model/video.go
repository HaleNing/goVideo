package model

import "github.com/jinzhu/gorm"

// Video 视频模型
type Video struct {
	gorm.Model
	TitleName string
	Info      string
}
