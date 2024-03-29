package config

import (
	"fmt"

	"github.com/pkg/errors"
	"go-micro.dev/v4/config"
	"go-micro.dev/v4/config/source/env"
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
	Port: 8666,
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
	configor, err := config.NewConfig(config.WithSource(env.NewSource()))
	if err != nil {
		return errors.Wrap(err, "configor.New")
	}
	if err := configor.Load(); err != nil {
		return errors.Wrap(err, "configor.Load")
	}
	if err := configor.Scan(cfg); err != nil {
		return errors.Wrap(err, "configor.Scan")
	}
	return nil
}
