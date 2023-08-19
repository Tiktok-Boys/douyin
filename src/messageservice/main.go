package main

import (
	"log"

	grpcc "github.com/go-micro/plugins/v4/client/grpc"
	grpcs "github.com/go-micro/plugins/v4/server/grpc"
	"go-micro.dev/v4"
	"go-micro.dev/v4/logger"

	"github.com/Tiktok-Boys/douyin/src/messageservice/config"
	"github.com/Tiktok-Boys/douyin/src/messageservice/dal"
	"github.com/Tiktok-Boys/douyin/src/messageservice/handler"
	pb "github.com/Tiktok-Boys/douyin/src/messageservice/proto"
)

var (
	name    = "messageservice"
	version = "1.0.0"
)

func main() {
	// Load conigurations
	if err := config.Load(); err != nil {
		logger.Fatal(err)
	}

	logger.Info(config.MySQL())

	// Init Data Access Layer
	dal.Init(config.MySQL())

	// Create service
	srv := micro.NewService(
		micro.Server(grpcs.NewServer()),
		micro.Client(grpcc.NewClient()),
	)
	opts := []micro.Option{
		micro.Name(name),
		micro.Version(version),
		micro.Address(config.Address()),
	}
	srv.Init(opts...)

	// Register handler
	if err := pb.RegisterMessageServiceHandler(srv.Server(), new(handler.MessageService)); err != nil {
		log.Fatal(err)
	}

	// Run service
	if err := srv.Run(); err != nil {
		logger.Fatal(err)
	}
}
