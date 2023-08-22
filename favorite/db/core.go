package db

import (
	"fmt"
	"github.com/redis/go-redis/v9"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var MysqlDB *gorm.DB
var RedisDB *redis.Client

func init() {
	//mysql_user := config.CFG.MySQL.Username
	//mysql_password := config.CFG.MySQL.Password
	//mysql_host := config.CFG.MySQL.Host
	//mysql_port := config.CFG.MySQL.Port
	//mysql_database := config.CFG.MySQL.Database
	mysql_user := "root"
	mysql_password := "tiktok"
	mysql_host := "0.0.0.0"
	mysql_port := "3306"
	mysql_database := "mysql"
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local", mysql_user, mysql_password, mysql_host, mysql_port, mysql_database)
	MysqlDB, _ = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	RedisDB = redis.NewClient(&redis.Options{
		Addr: "0.0.0.0:6379",
		DB:   0,
	})
}
