package handler

import (
	"net/http"
	"publish/dal/db"
	"time"

	"github.com/gin-gonic/gin"
)

func Publish(ctx *gin.Context) error {
	//TODO: 这一步应该用rpc获取
	title := ctx.PostForm("title")
	// token := ctx.PostForm("token")
	// fileHeader, _ := ctx.FormFile("data")

	//生成视频地址

	//生成图片地址

	//开启协程上传

	//通过token获取userId
	//TODO: 问问怎么通过token获取userId
	var userId int64
	userId = 1

	//构造video
	video := &db.Video{
		Id:            userId,
		Title:         title,
		Cover_url:     "",
		Play_url:      "",
		FavoriteCount: 0,
		CommentCount:  0,
		CreateAt:      time.Now(),
	}

	//上传video到数据库
	if _, err := db.CreateVideo(video); err != nil {
		ctx.JSON(http.StatusOK, gin.H{"res": "Failed"})
		return err
	}

	//TODO: 这一步应该用rpc返回response
	ctx.JSON(http.StatusOK, gin.H{"res": "OK"})
}
