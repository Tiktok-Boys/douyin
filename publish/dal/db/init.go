package db

import (
	"fmt"

	"go-micro.dev/v4/logger"

	"publish/config"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Init(cfg config.MySQLConfig) {
	var err error
	mysql_user := cfg.Username
	mysql_password := cfg.Password
	mysql_host := cfg.Host
	mysql_port := cfg.Port
	mysql_database := cfg.Database
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local", mysql_user, mysql_password, mysql_host, mysql_port, mysql_database)
	logger.Info(dsn)
	DB, err = gorm.Open(mysql.Open(dsn),
		&gorm.Config{
			PrepareStmt:            true,
			SkipDefaultTransaction: true,
		},
	)
	if err != nil {
		panic(err)
	}
}
