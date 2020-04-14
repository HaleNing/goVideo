package serializer

import "goVideo/model"

// 视频序列化器
type Video struct {
	ID        uint   `json:"id"`
	TitleName string `json:"title_name"`
	Info      string `json:"info"`
	CreatedAt int64  `json:"created_at"`
}

// BuildVideo 序列化视频
func BuildVideo(video model.Video) Video {
	return Video{
		ID:        video.ID,
		TitleName: video.TitleName,
		Info:      video.Info,
		CreatedAt: video.CreatedAt.Unix(),
	}
}

// BuildVideos 序列化视频列表
func BuildVideos(items []model.Video) []Video {
	var videos []Video
	// 循环调用BuildVideo方法
	for _, item := range items {
		video := BuildVideo(item)
		videos = append(videos, video)
	}
	return videos
}
