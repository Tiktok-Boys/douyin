package main

import (
	"publish/handler"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.POST("/douyin/publish_2/action", handler.Publish)
}
