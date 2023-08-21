package dal

import (
	"feed/config"
	"feed/dal/db"
)

// init dal
func Init(mysql config.MySQLConfig) {
	db.Init(mysql)
}
