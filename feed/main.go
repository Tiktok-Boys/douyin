package main

import (
	"github.com/Tiktok-Boys/douyin/feed/config"
	"github.com/Tiktok-Boys/douyin/feed/dal"
	"github.com/Tiktok-Boys/douyin/feed/handler"
	"github.com/gin-gonic/gin"
	"log"
)

func main() {
	// Load conigurations
	if err := config.Load(); err != nil {
		logger.Fatal(err)
	}

	logger.Info(config.MySQL())

	// Init Data Access Layer
	dal.Init(config.MySQL())

	router := gin.Default()

	router.GET("/douyin/feed",handler.Feed)

	router.run(:9999)
}
