package handler

import (
	"context"
	proto "github.com/Tiktok-Boys/douyin/api/proto/favorite"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

var (
	FavoriteServiceClient proto.FavoriteService
)

func FavoriteAction(ginCtx *gin.Context) {
	actionType, _ := strconv.Atoi(ginCtx.Query("action_type"))
	token := ginCtx.Query("token")
	vid, _ := strconv.ParseInt(ginCtx.Query("video_id"), 10, 64)
	response, err := FavoriteServiceClient.FavoriteAction(context.TODO(), &proto.FavoriteActionRequest{
		Token:      token,
		VideoId:    vid,
		ActionType: int32(actionType),
	})
	if err != nil {
		ginCtx.JSON(http.StatusInternalServerError, gin.H{
			"status_code": 500,
			"status_msg":  "FavoriteService.FavoriteAction: " + err.Error(),
		})
		return
	}
	ginCtx.JSON(http.StatusOK, gin.H{
		"status_code": response.StatusCode,
		"status_msg":  response.StatusMsg,
	})
}

func FavoriteList(ginCtx *gin.Context) {
	token := ginCtx.Query("token")
	uid, _ := strconv.ParseInt(ginCtx.Query("user_id"), 10, 64)
	response, err := FavoriteServiceClient.FavoriteList(context.TODO(), &proto.FavoriteListRequest{
		UsrId: uid,
		Token: token,
	})
	if err != nil {
		ginCtx.JSON(http.StatusInternalServerError, proto.FavoriteListResponse{
			StatusCode: 500,
			StatusMsg:  err.Error(),
			VideoList:  nil,
		})
		return
	}
	ginCtx.JSON(http.StatusOK, proto.FavoriteListResponse{
		StatusCode: response.StatusCode,
		StatusMsg:  response.StatusMsg,
		VideoList:  response.VideoList,
	})
}
