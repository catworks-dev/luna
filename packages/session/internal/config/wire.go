//go:build wireinject

package config

import (
	"catworks/luna/session/internal/domain"
	"catworks/luna/session/internal/repository"
	"github.com/google/wire"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type Container struct {
	Config         *Config
	Logger         *logrus.Logger
	DB             *gorm.DB
	SessionStorage domain.SessionStorage
}

func NewContainer(cfg *Config) (*Container, error) {
	panic(
		wire.Build(
			NewLogger,
			NewGorm,
			repository.NewSessionRepository,
			wire.Struct(new(Container), "*"),
		),
	)
}
