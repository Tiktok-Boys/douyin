package db

import (
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"github.com/redis/go-redis/v9"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var MysqlDB *gorm.DB
var RedisDB0 *redis.Client // token
var MinioClient *minio.Client
var err error

func init() {

	MysqlDB, err = gorm.Open(mysql.Open("tiktok:TikTokBoys123@tcp(kasperxms.xyz:3306)/tiktok?charset=utf8mb4&parseTime=True&loc=Local"))
	if err != nil {
		panic(err)
	}

	RedisDB0 = redis.NewClient(&redis.Options{
		Addr:     "kasperxms.xyz:6379",
		Password: "TikTokBoys123",
		DB:       0,
	})

	endpoint := "kasperxms.xyz:9000"
	accessKeyID := "IiGUOYoJPQntPHuipHSx"
	secretAccessKey := "fMpdpvGqEtHs5KKAdyugpZyIRi674X4PD0Y0zSJy"

	// 初始化一个minio客户端对象
	MinioClient, err = minio.New(endpoint, &minio.Options{
		Creds: credentials.NewStaticV4(accessKeyID, secretAccessKey, ""),
	})
}
