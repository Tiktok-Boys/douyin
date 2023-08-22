package main

import (
	"go-micro.dev/v4"
	"go-micro.dev/v4/logger"
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
		// micro.Registry(etcd.NewRegistry(registry.Addrs(config.EtcdAddress()))),
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
