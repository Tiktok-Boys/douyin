package main

import (
	"publish/handler"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.POST("/douyin/publish/action", handler.Publish)
}
