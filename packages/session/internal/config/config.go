package config

import (
	"github.com/ilyakaznacheev/cleanenv"
	"os"
	"time"
)

type Config struct {
	StoragePath string        `yaml:"storage_path" env-required:"true"`
	SessionTTL  time.Duration `yaml:"session_ttl" env-default:"2160h"`
	Grpc        GrpcConfig    `yaml:"grpc"`
}

type GrpcConfig struct {
	Port    int           `yaml:"port"`
	Timeout time.Duration `yaml:"timeout"`
}

func Require(path string) *Config {
	if path == "" {
		panic("Config path must not be empty")
	}

	if _, err := os.Stat(path); os.IsNotExist(err) {
		panic("Config file not found")
	}

	var cfg Config
	if err := cleanenv.ReadConfig(path, &cfg); err != nil {
		panic("Unable to read config file: " + err.Error())
	}

	return &cfg
}
