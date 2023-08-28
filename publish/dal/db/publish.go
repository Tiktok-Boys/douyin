package db

import (
	"time"
)

type Video struct {
	Id         int64
	User_id    int64
	Play_url   string
	Cover_url  string
	Title      string
	Created_at time.Time
	Updated_at time.Time
	Deleted_at time.Time
}

func (Video) TableName() string {
	return "video"
}

// 存入一个新的Video，返回Video实例
func CreateVideo(video *Video) (*Video, error) {
	/*Video := Video{Name: Videoname, Password: password, FollowingCount: 0, FollowerCount: 0, CreateAt: time.Now()}*/

	result := DB.Create(&video)

	if result.Error != nil {
		return nil, result.Error
	}

	return video, nil
}

// 根据videoid，查找video实体
func FindVideoById(videoId int64) (*Video, error) {
	video := Video{Id: videoId}

	result := DB.Where("video_id = ?", videoId).First(&video)
	err := result.Error
	if err != nil {
		return nil, err
	}
	return &video, err
}
