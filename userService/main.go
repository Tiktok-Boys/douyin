package main

import (
	"github.com/go-micro/plugins/v4/registry/etcd"
	grpcs "github.com/go-micro/plugins/v4/server/grpc"
	"go-micro.dev/v4"
	"go-micro.dev/v4/logger"
	"go-micro.dev/v4/registry"
	"userService/config"
	"userService/handler"
	pb "userService/proto"
)

var (
	service = "tiktokboys.douyin.user"
	version = "latest"
)

func main() {
	if err := config.Load(); err != nil {
		logger.Fatal(err)
	}

	logger.Info(config.EtcdAddress())

	// Create service
	srv := micro.NewService(
		micro.Server(grpcs.NewServer()),
	)
	srv.Init(
		micro.Name(service),
		micro.Version(version),
		micro.Address(config.Address()),
		micro.Registry(etcd.NewRegistry(registry.Addrs(config.EtcdAddress()))),
	)

	// Register handler
	if err := pb.RegisterUserServiceHandler(srv.Server(), handler.NewUserService()); err != nil {
		logger.Fatal(err)
	}
	// Run service
	if err := srv.Run(); err != nil {
		logger.Fatal(err)
	}
}
