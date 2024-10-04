//go:build wireinject

package config

import (
	"github.com/google/wire"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type Container struct {
	Config *Config
	Logger *logrus.Logger
	DB     *gorm.DB
}

func NewContainer(cfg *Config) (*Container, error) {
	panic(
		wire.Build(
			NewLogger,
			NewGorm,
			wire.Struct(new(Container), "*"),
		),
	)
}
