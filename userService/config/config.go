package config

import "fmt"

type Config struct {
	Etcd EtcdConfig
}

var cfg *Config = &Config{}

type EtcdConfig struct {
	Host string
	Port int
}

func EtcdAddress() string {
	return fmt.Sprintf("%s:%d", cfg.Etcd.Host, cfg.Etcd.Port)
}
