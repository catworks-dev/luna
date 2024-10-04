package config

import (
	"github.com/ilyakaznacheev/cleanenv"
	"github.com/sirupsen/logrus"
	"os"
	"time"
)

type Config struct {
	StoragePath string        `yaml:"storage_path" env-required:"true"`
	SessionTTL  time.Duration `yaml:"session_ttl" env-default:"2160h"`
	Grpc        GrpcConfig    `yaml:"grpc"`
	LogLevel    string        `yaml:"log_level" env-default:"info"`
	Version     string        `yaml:"version" env-default:"v0.1.0"`
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

func NewLogger(config *Config) *logrus.Logger {
	logger := logrus.New()
	logger.SetFormatter(&logrus.TextFormatter{})
	logger.SetOutput(os.Stdout)

	if lvl, err := logrus.ParseLevel(config.LogLevel); err != nil {
		logger.WithField("log_level", config.LogLevel).Warn("Invalid log level, using info")
		logger.SetLevel(logrus.InfoLevel)
	} else {
		logger.SetLevel(lvl)
	}

	return logger
}
