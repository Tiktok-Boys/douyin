package main

import (
	"favorite/handler"
	pb "favorite/proto"

	"go-micro.dev/v4"
	"go-micro.dev/v4/logger"
)

var (
	service = "favorite"
	version = "latest"
)

func main() {
	// Create service
	srv := micro.NewService()
	srv.Init(
		micro.Name(service),
		micro.Address("127.0.0.1:8866"),
		micro.Version(version),
	)

	// Register handler
	if err := pb.RegisterFavoriteServiceHandler(srv.Server(), handler.NewFavorite()); err != nil {
		logger.Fatal(err)
	}
	// Run service
	if err := srv.Run(); err != nil {
		logger.Fatal(err)
	}
}
