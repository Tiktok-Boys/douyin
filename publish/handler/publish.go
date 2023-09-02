package handler

import (
	"fmt"
	"net/http"
	"path/filepath"
	"time"

	"github.com/gin-gonic/gin"

	"publish/dal/db"
	minio_api "publish/minio-api"
)

func Publish(ctx *gin.Context) {
	//TODO: 这一步应该用rpc获取
	title := ctx.PostForm("title")
	fmt.Println("title = ", title)
	// token := ctx.PostForm("token")
	// fileHeader, _ := ctx.FormFile("data")

	//通过token获取userId
	//TODO: 问问怎么通过token获取userId
	var userId int64
	userId = 2

	//生成视频地址
	data, err := ctx.FormFile("data")
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"status_code": -1,
			"status_msg":  "can not get file",
		})
		return
	}

	filename := filepath.Base(data.Filename)
	finalName := fmt.Sprintf("%d_%s", userId, filename)
	saveFile := filepath.Join("./public/", finalName)
	if err := ctx.SaveUploadedFile(data, saveFile); err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"status_code": -1,
			"status_msg":  "can not update this file",
		})
		return
	}

	client := minio_api.InitMinioClient()
	minio_api.PutObjects(client, "video", "./public", finalName)
	//生成图片地址

	//开启协程上传

	//构造video
	video := &db.Video{
		Id:         0,
		User_id:    userId,
		Title:      title,
		Cover_url:  "",
		Play_url:   saveFile,
		Created_at: time.Now(),
	}
	fmt.Println("video ", video)
	//上传video到数据库
	if _, err := db.CreateVideo(video); err != nil {
		ctx.JSON(http.StatusOK, gin.H{"res": "Failed"})
		return
	}

	//TODO: 这一步应该用rpc返回response
	ctx.JSON(http.StatusOK, gin.H{"res": "OK"})
	return
}
