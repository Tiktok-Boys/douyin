package handler

import (
	"context"
	proto "favorite/proto"
)

type Favorite struct {
	auth   *Authentication
	mapper *Mapper
}

func NewFavorite() *Favorite {
	return &Favorite{
		auth:   NewAuthentication(),
		mapper: NewMapper(),
	}
}

func (f *Favorite) FavoriteAction(ctx context.Context, request *proto.FavoriteActionRequest, response *proto.FavoriteActionResponse) error {
	actionType := request.ActionType //1点赞  2取消
	vid := request.VideoId
	token := request.Token
	// 鉴权
	usrId, err := f.auth.ValidateToken(ctx, token)
	if err != nil {
		response.StatusCode = 500
		response.StatusMsg = "鉴权失败"
		return err
	}
	err = f.mapper.FavoriteAction(ctx, usrId, vid, actionType, token)
	if err != nil {
		response.StatusCode = 500
		response.StatusMsg = "点赞失败"
		return err
	}
	response.StatusMsg = "successful"
	response.StatusCode = 0
	return nil
}

func (f *Favorite) FavoriteList(ctx context.Context, request *proto.FavoriteListRequest, response *proto.FavoriteListResponse) error {
	usrId := request.UsrId
	token := request.Token
	viedoList, err := f.mapper.FavoriteList(ctx, usrId, token)
	if err != nil {

		return err
	}
	response.VideoList = viedoList
	response.StatusCode = 0
	response.StatusMsg = "succeed"
	return nil
}
