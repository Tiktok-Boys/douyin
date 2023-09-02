package dal

import (
	"publish/config"
	"publish/dal/db"
)

// init dal
func Init(mysql config.MySQLConfig) {
	db.Init(mysql)
}
