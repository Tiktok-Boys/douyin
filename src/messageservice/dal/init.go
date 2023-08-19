package dal

import (
	"github.com/Tiktok-Boys/douyin/src/messageservice/config"
	"github.com/Tiktok-Boys/douyin/src/messageservice/dal/db"
)

// Init init dal
func Init(mysql config.MySQLConfig) {
	db.Init(mysql) // mysql
}
