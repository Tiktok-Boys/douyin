package db

import (
	"github.com/redis/go-redis/v9"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var MysqlDB *gorm.DB
var RedisDB *redis.Client

func init() {
	MysqlDB, _ = gorm.Open(mysql.Open("tiktok:TikTokBoys123@tcp(kasperxms.xyz:3306)/tiktok?charset=utf8mb4&parseTime=True&loc=Local"), &gorm.Config{})
	RedisDB = redis.NewClient(&redis.Options{
		Addr:     "",
		Password: "",
		DB:       0,
	})
}
