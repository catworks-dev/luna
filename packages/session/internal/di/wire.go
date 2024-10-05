//go:build wireinject

package di

import (
	"catworks/luna/session/internal/config"
	"catworks/luna/session/internal/domain"
	"catworks/luna/session/internal/repository"
	"catworks/luna/session/internal/service"
	"catworks/luna/session/internal/transport/rpc"
	"catworks/luna/session/internal/usecase"
	"github.com/google/wire"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type Container struct {
	Config         *config.Config
	Logger         *logrus.Logger
	DB             *gorm.DB
	SessionStorage domain.SessionStorage
	Server         *rpc.Server
}

func NewContainer(cfg *config.Config) (*Container, error) {
	panic(
		wire.Build(
			config.NewLogger,
			config.NewGorm,
			repository.NewSessionRepository,
			service.NewJWTService,
			usecase.NewSessionUseCase,
			rpc.NewServer,
			wire.Struct(new(Container), "*"),
		),
	)
}
