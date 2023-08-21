package dal

import (
	"github.com/Tiktok-Boys/douyin/feed/config"
	"github.com/Tiktok-Boys/douyin/feed/dal/db"
)

// init dal
func Init(mysql config.MySQLConfig) {
	db.Init(mysql)
}
