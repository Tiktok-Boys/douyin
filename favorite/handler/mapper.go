package handler

import (
	"context"
	"errors"
	"favorite/db"
	"favorite/model"
	favorite "favorite/proto"
	"gorm.io/gorm"
	"time"
)

type Mapper struct {
	db *gorm.DB
}

func NewMapper() *Mapper {
	return &Mapper{db: db.MysqlDB}
}

func (m *Mapper) FavoriteAction(ctx context.Context, uid, vid int64, actionType int32, token string) error {
	var fav model.Like
	err := m.db.Where("video_id = ?", vid).First(&fav, uid).Error
	if !errors.Is(err, gorm.ErrRecordNotFound) {
		m.db.Where("video_id = ?", vid).Delete(&model.Like{}, uid)
	}
	like := model.Like{
		User_id:    uid,
		Video_id:   vid,
		Created_at: time.Now(),
	}
	m.db.Create(&like)
	return nil
}

func (m *Mapper) FavoriteList(ctx context.Context, uid int64, token string) ([]*favorite.Video, error) {
	var videoList []*favorite.Video
	var queryRes []model.Like
	// 查到点赞过的视频
	m.db.Model(&favorite.Video{}).Where("user_id = ?", uid).Find(&queryRes)
	for _, v := range queryRes {
		// 根据点赞的视频找到具体的视频信息
		var modelVideo model.Video
		m.db.First(&modelVideo, v.Video_id)
		// 找到视频的作者信息
		var user model.User
		m.db.First(&user, modelVideo.User_id)
		video := &favorite.Video{
			Id: v.Video_id,
			Author: &favorite.User{
				Id:              user.Id,
				Name:            user.Name,
				FollowCount:     0,
				FollowerCount:   0,
				IsFollow:        false,
				Avatar:          user.Avatar,
				BackgroundImage: user.Background_image,
				Signature:       user.Signature,
				TotalFavorited:  "",
				WorkCount:       0,
				FavoriteCount:   0,
			},
			PlayUrl:       modelVideo.Play_url,
			CoverUrl:      modelVideo.Cover_url,
			FavoriteCount: 0,
			CommentCount:  0,
			IsFavorite:    true,
			Title:         modelVideo.Title,
		}
		videoList = append(videoList, video)
	}
	return []*favorite.Video{}, nil
}
