package serializer

import "goVideo/model"

// 视频序列化器
type Video struct {
	ID        uint   `json:"id"`
	TitleName string `json:"title_name"`
	Info      string `json:"info"`
	Url       string `json:"url"`
	View      uint64 `json:"view"`
	User      User   `json:"user"`
	CreatedAt int64  `json:"created_at"`
}

// BuildVideo 序列化视频
func BuildVideo(video model.Video) Video {
	user, _ := model.GetUser(video.UserId)

	return Video{
		ID:        video.ID,
		TitleName: video.TitleName,
		Info:      video.Info,
		Url:       video.Url,
		View:      video.View(),
		User:      BuildUser(user),
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
