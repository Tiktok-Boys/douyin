package config

import (
	"fmt"

	"github.com/pkg/errors"
	"go-micro.dev/v4/config"
	"go-micro.dev/v4/config/source/env"
)

type Config struct {
	MySQL MySQLConfig
	Etcd  EtcdConfig
}

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

var CFG *Config = &Config{
	// Port: 9777,
}

func MySQL() MySQLConfig {
	return CFG.MySQL
}

func EtcdAddress() string {
	return fmt.Sprintf("%s:%d", CFG.Etcd.Host, CFG.Etcd.Port)
}

func Load() error {
	configor, err := config.NewConfig(config.WithSource(env.NewSource()))
	if err != nil {
		return errors.Wrap(err, "configor.New")
	}
	if err := configor.Load(); err != nil {
		return errors.Wrap(err, "configor.Load")
	}
	if err := configor.Scan(CFG); err != nil {
		return errors.Wrap(err, "configor.Scan")
	}
	return nil
}
