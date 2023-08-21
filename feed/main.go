package main

import (
	"feed/config"
	"feed/dal"
	"go-micro.dev/v4/logger"
)

func main() {
	// Load conigurations
	if err := config.Load(); err != nil {
		logger.Fatal(err)
	}

	logger.Info(config.MySQL())

	// Init Data Access Layer
	dal.Init(config.MySQL())

}
