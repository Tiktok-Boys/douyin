package main

import (
	"github.com/Tiktok-Boys/douyin/api/config"
	"github.com/Tiktok-Boys/douyin/api/handler"
	"github.com/Tiktok-Boys/douyin/api/middleware"
	favorite "github.com/Tiktok-Boys/douyin/api/proto/favorite"
	message "github.com/Tiktok-Boys/douyin/api/proto/message"
	user "github.com/Tiktok-Boys/douyin/api/proto/user"
	"github.com/gin-gonic/gin"

	grpcc "github.com/go-micro/plugins/v4/client/grpc"
	"github.com/go-micro/plugins/v4/registry/etcd"
	"go-micro.dev/v4"
	"go-micro.dev/v4/registry"
)

var (
	messageServiceName  = "tiktokboys.douyin.message"
	favoriteServiceName = "tiktokboys.douyin.favorite"
	userServiceName     = "tiktokboys.douyin.user"
)

func main() {
	// Load conigurations
	if err := config.Load(); err != nil {
		panic(err)
	}

	serviceDiscovery(config.EtcdAddress())

	r := setupRouter()
	// Listen and Server in 0.0.0.0:8080
	r.Run(":8080")
}

func setupRouter() *gin.Engine {
	r := gin.Default()

	// Logger middleware will write the logs to gin.DefaultWriter even if you set with GIN_MODE=release.
	// By default gin.DefaultWriter = os.Stdout
	r.Use(gin.Logger())

	// Recovery middleware recovers from any panics and writes a 500 if there was one.
	r.Use(gin.Recovery())

	// Routes that do not require login.
	// r.POST("/account/register", handler.RegisterHandler)

	// Routes that need to be authenticated.
	authenticated := r.Group("/")
	authenticated.Use(middleware.Authenticate())
	{
		// Message service
		message := authenticated.Group("/douyin/message")
		message.GET("/chat", handler.GetChatHandler)
		message.POST("/action", handler.ActMessageHandler)

		// favorite service
		favorite := authenticated.Group("/douyin/favorite")
		{
			favorite.POST("/action", handler.FavoriteAction)
			favorite.GET("/list", handler.FavoriteList)
		}

		// user service
		user := authenticated.Group("/douyin/user")
		{
			user.POST("/login", handler.LoginHandler)
			user.POST("/register", handler.RegisterHandler)
			user.GET("/", handler.UserInfoHandler)
		}
	}

	return r
}

func serviceDiscovery(etcdAddr string) {
	etcdReg := etcd.NewRegistry(
		registry.Addrs(etcdAddr),
	)

	messageService := micro.NewService(
		micro.Client(grpcc.NewClient()),
	)
	messageService.Init(micro.Registry(etcdReg))
	handler.MessageServiceClient = message.NewMessageService(messageServiceName, messageService.Client())

	favoriteService := micro.NewService(
		micro.Client(grpcc.NewClient()),
	)
	favoriteService.Init(micro.Registry(etcdReg))
	handler.FavoriteServiceClient = favorite.NewFavoriteService(favoriteServiceName, favoriteService.Client())

	userService := micro.NewService(
		micro.Client(grpcc.NewClient()),
	)
	handler.UserServiceClient = user.NewUserService(userServiceName, userService.Client())
}
