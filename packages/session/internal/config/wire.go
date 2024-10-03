//go:build wireinject

package config

import (
	"github.com/google/wire"
)

type Container struct {
	Config *Config
	//Logger *logrus.Logger
}

func NewContainer(cfg *Config) (*Container, error) {
	panic(
		wire.Build(
			wire.Struct(new(Container), "*"),
		),
	)
}
