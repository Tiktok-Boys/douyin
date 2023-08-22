package handler

import (
	"context"
	proto "github.com/Tiktok-Boys/douyin/api/proto/user"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

var (
	UserServiceClient proto.UserService
)

func RegisterHandler(ctx *gin.Context) {
	username := ctx.Query("username")
	password := ctx.Query("password")
	rsp, err := UserServiceClient.Register(context.TODO(), &proto.DouyinUserRegisterRequest{
		Username: username,
		Password: password,
	})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"status_code": 500,
			"status_msg":  "UserService.Register: " + err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"status_code": rsp.StatusCode,
		"status_msg":  rsp.StatusMsg,
		"user_id":     rsp.UserId,
		"token":       rsp.Token,
	})
}
func LoginHandler(ctx *gin.Context) {
	username := ctx.Query("username")
	password := ctx.Query("password")
	rsp, err := UserServiceClient.Login(context.TODO(), &proto.DouyinUserLoginRequest{
		Username: username,
		Password: password,
	})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"status_code": 500,
			"status_msg":  "UserService.Register: " + err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"status_code": rsp.StatusCode,
		"status_msg":  rsp.StatusMsg,
		"user_id":     rsp.UserId,
		"token":       rsp.Token,
	})
}
func UserInfoHandler(ctx *gin.Context) {
	token := ctx.Query("token")
	userId, _ := strconv.ParseInt(ctx.Query("user_id"), 10, 64)
	rsp, err := UserServiceClient.UserInfo(context.TODO(), &proto.DouyinUserRequest{
		Token:  token,
		UserId: userId,
	})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"status_code": 500,
			"status_msg":  "UserService.Register: " + err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"status_code": rsp.StatusCode,
		"status_msg":  rsp.StatusMsg,
		"user":        rsp.User,
	})
}
