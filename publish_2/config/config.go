package config

import (
	"fmt"
	"github.com/pkg/errors"
	"go-micro.dev/v4/config"
	"go-micro.dev/v4/config/source/env"
)

type Config struct {
	Port int64
	Etcd EtcdConfig
}

var cfg *Config = &Config{}

type EtcdConfig struct {
	Host string
	Port int
}

func Address() string {
	return fmt.Sprintf("127.0.0.1:%d", cfg.Port)
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
