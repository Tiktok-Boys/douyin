package main

import (
	"log"

	grpcc "github.com/go-micro/plugins/v4/client/grpc"
	"github.com/go-micro/plugins/v4/registry/etcd"
	grpcs "github.com/go-micro/plugins/v4/server/grpc"
	"go-micro.dev/v4"
	"go-micro.dev/v4/logger"
	"go-micro.dev/v4/registry"

	"github.com/Tiktok-Boys/douyin/messageservice/config"
	"github.com/Tiktok-Boys/douyin/messageservice/dal"
	"github.com/Tiktok-Boys/douyin/messageservice/handler"
	pb "github.com/Tiktok-Boys/douyin/messageservice/proto"
)

var (
	name    = "tiktokboys.douyin.message"
	version = "1.0.0"
)

func main() {
	// Load conigurations
	if err := config.Load(); err != nil {
		logger.Fatal(err)
	}

	logger.Info(config.MySQL())
	logger.Info(config.EtcdAddress())

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
		micro.Registry(etcd.NewRegistry(
			registry.Addrs(config.EtcdAddress()),
		)),
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
