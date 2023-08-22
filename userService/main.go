package main

import (
	"github.com/go-micro/plugins/v4/registry/etcd"
	"go-micro.dev/v4"
	"go-micro.dev/v4/logger"
	"go-micro.dev/v4/registry"
	"userService/config"
	"userService/handler"
	pb "userService/proto"
)

var (
	service = "userservice"
	version = "latest"
)

func main() {
	// Create service
	srv := micro.NewService()
	srv.Init(
		micro.Name(service),
		micro.Version(version),
		micro.Registry(etcd.NewRegistry(registry.Addrs(config.EtcdAddress()))),
	)

	// Register handler
	if err := pb.RegisterUserServiceHandler(srv.Server(), new(handler.UserService)); err != nil {
		logger.Fatal(err)
	}
	// Run service
	if err := srv.Run(); err != nil {
		logger.Fatal(err)
	}
}
