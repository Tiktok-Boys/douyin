package config

import (
	"fmt"
)

type Config struct {
	Port  int
	MySQL MySQLConfig
	Etcd  EtcdConfig
}

// type RedisConfig struct {
// 	Addr string
// }

// type TracingConfig struct {
// 	Enable bool
// 	Jaeger JaegerConfig
// }

// type JaegerConfig struct {
// 	URL string
// }

type MySQLConfig struct {
	Host     string
	Port     int
	Username string
	Password string
	Database string
}

type EtcdConfig struct {
	Host string
	Port int
}

var cfg *Config = &Config{
	Port: 8999,
}

func Address() string {
	return fmt.Sprintf(":%d", cfg.Port)
}

// func Redis() RedisConfig {
// 	return cfg.Redis
// }

// func Tracing() TracingConfig {
// 	return cfg.Tracing
// }

func MySQL() MySQLConfig {
	return cfg.MySQL
}

func EtcdAddress() string {
	return fmt.Sprintf("%s:%d", cfg.Etcd.Host, cfg.Etcd.Port)
}

func Load() error {
	cfg.MySQL.Host = "47.100.235.108"
	cfg.MySQL.Port = 13306
	cfg.MySQL.Username = "root"
	cfg.MySQL.Password = "root"
	cfg.MySQL.Database = "tiktok"

	return nil
}
