package main

import (
	"feed/config"
	"feed/dal"
	"feed/handler"
	pb "feed/proto"
	"github.com/go-micro/plugins/v4/registry/etcd"
	"go-micro.dev/v4"
	"go-micro.dev/v4/logger"
	"go-micro.dev/v4/registry"
)

var (
	service = "feed"
	version = "latest"
)

func main() {
	// Load conigurations
	if err := config.Load(); err != nil {
		logger.Fatal(err)
	}

	logger.Info(config.MySQL())

	// Init Data Access Layer
	dal.Init(config.MySQL())

	srv := micro.NewService()
	srv.Init(
		micro.Name(service),
		micro.Address("127.0.0.1:8866"),
		micro.Version(version),
		micro.Registry(etcd.NewRegistry(
			registry.Addrs(config.EtcdAddress()),
		)),
	)

	// Register handler
	if err := pb.RegisterFeedServiceHandler(srv.Server(), new(handler.FeedService)); err != nil {
		logger.Fatal(err)
	}
	// Run service
	if err := srv.Run(); err != nil {
		logger.Fatal(err)
	}
}
