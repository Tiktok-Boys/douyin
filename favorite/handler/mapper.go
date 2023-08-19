package handler

import (
	"context"
	"favorite/db"
	favorite "favorite/proto"
	"gorm.io/gorm"
)

type Mapper struct {
	db *gorm.DB
}

func NewMapper() *Mapper {
	return &Mapper{db: db.MysqlDB}
}

func (m *Mapper) FavoriteAction(ctx context.Context, uid, vid int64, actionType int32, token string) error {

	return nil
}

func (m *Mapper) FavoriteList(ctx context.Context, uid int64, token string) ([]*favorite.Video, error) {

	return []*favorite.Video{}, nil
}
