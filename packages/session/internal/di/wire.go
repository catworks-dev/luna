//go:build wireinject

package di

import (
	"catworks/luna/session/internal/config"
	"catworks/luna/session/internal/repository"
	"catworks/luna/session/internal/service"
	"catworks/luna/session/internal/transport/rpc"
	"catworks/luna/session/internal/usecase"
	"github.com/google/wire"
)

func NewServer(cfg *config.Config) (*rpc.Server, error) {
	panic(
		wire.Build(
			config.NewLogger,
			wire.NewSet(
				wire.Struct(new(config.GormOptions), "*"),
				config.NewGorm,
			),
			repository.NewSessionRepository,
			service.NewJWTService,
			wire.NewSet(
				wire.Struct(new(usecase.SessionUseCaseOptions), "*"),
				usecase.NewSessionUseCase,
			),
			wire.NewSet(
				wire.Struct(new(rpc.ServerOptions), "*"),
				rpc.NewServer,
			),
		),
	)
}
