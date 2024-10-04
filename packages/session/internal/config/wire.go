//go:build wireinject

package config

import (
	"github.com/google/wire"
	"github.com/sirupsen/logrus"
)

type Container struct {
	Config *Config
	Logger *logrus.Logger
}

func NewContainer(cfg *Config) (*Container, error) {
	panic(
		wire.Build(
			NewLogger,
			wire.Struct(new(Container), "*"),
		),
	)
}
