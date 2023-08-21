package db

import (
	"fmt"
	"sync"
	"time"
)

type Video struct {
	VideoId       int64  `gorm:"primaryKey"`
	UserId        int64  `gorm:"default:(-)"`
	PlayUrl       string `gorm:"default:(-)"`
	CoverUrl      string `gorm:"default:(-)"`
	FavoriteCount int64  `gorm:"default:(-)"`
	CommentCount  int64  `gorm:"default:(-)"`
	Title         string `gorm:"default:(-)"`
	CreateAt      time.Time
	UpdateAt      time.Time
	DeleteAt      time.Time
}
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

// 根据videoid，查找video实体
func FindVideoById(id int64) (*Video, error) {
	video := Video{Id: id}

	result := DB.Where("Video_id = ?", id).First(&video)
	err := result.Error
	if err != nil {
		return nil, err
	}
	return &video, err
}

// 根据UserId，查出Video列表
func QueryVideoByUserId(userId int64) ([]*Video, error) {
	var videos []*Video
	err := DB.Where("user_id = ?", userId).Find(&videos).Error
	if err != nil {
		return nil, err
	}
	return videos, nil
}

// 根据时间和需要查询的条数，获取video列表
func (*VideoDao) QueryVideo(date *string, limit int) []*Video {
	var VideoList []*Video
	DB.Where("create_at < ?", *date).Order("create_at desc").Find(&VideoList)
	if len(VideoList) <= limit {
		return VideoList
	}
	return VideoList[0:limit]
}
