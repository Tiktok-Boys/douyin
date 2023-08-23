package db

import (
	"github.com/redis/go-redis/v9"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var MysqlDB *gorm.DB
var RedisDB0 *redis.Client // token
var RedisDB1 *redis.Client // password
var RedisDB2 *redis.Client // following count
var RedisDB3 *redis.Client // follower count
var RedisDB4 *redis.Client // favorited count
var RedisDB5 *redis.Client // work count
var RedisDB6 *redis.Client // favs count
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

	RedisDB1 = redis.NewClient(&redis.Options{
		Addr:     "kasperxms.xyz:6379",
		Password: "TikTokBoys123",
		DB:       1,
	})

	RedisDB2 = redis.NewClient(&redis.Options{
		Addr:     "kasperxms.xyz:6379",
		Password: "TikTokBoys123",
		DB:       2,
	})

	RedisDB3 = redis.NewClient(&redis.Options{
		Addr:     "kasperxms.xyz:6379",
		Password: "TikTokBoys123",
		DB:       3,
	})

	RedisDB4 = redis.NewClient(&redis.Options{
		Addr:     "kasperxms.xyz:6379",
		Password: "TikTokBoys123",
		DB:       4,
	})

	RedisDB5 = redis.NewClient(&redis.Options{
		Addr:     "kasperxms.xyz:6379",
		Password: "TikTokBoys123",
		DB:       5,
	})

	RedisDB6 = redis.NewClient(&redis.Options{
		Addr:     "kasperxms.xyz:6379",
		Password: "TikTokBoys123",
		DB:       6,
	})
}
