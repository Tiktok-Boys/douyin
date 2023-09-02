package main

import (
	"github.com/gin-gonic/gin"
	"go-micro.dev/v4/logger"

	"publish/config"
	"publish/dal"
	"publish/handler"
)

func main() {
	if err := config.Load(); err != nil {
		logger.Fatal(err)
	}

	logger.Info(config.MySQL())

	// Init Data Access Layer
	dal.Init(config.MySQL())

	router := gin.Default()
	router.POST("/douyin/publish/action", handler.Publish)
	router.Run(":8888")
}
