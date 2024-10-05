package config

import (
	"catworks/luna/session/internal/repository"
	gl "github.com/onrik/gorm-logrus"
	"github.com/sirupsen/logrus"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var models = []interface{}{
	&repository.Session{},
}

type GormOptions struct {
	Config *Config
	Logger *logrus.Logger
}

func NewGorm(opts *GormOptions) *gorm.DB {
	gormConfig := &gorm.Config{}
	if opts.Config.LogDB {
		l := gl.New()
		l.SkipErrRecordNotFound = true
		gormConfig.Logger = l
	}
	db, err := gorm.Open(
		sqlite.Open(opts.Config.StoragePath),
		gormConfig,
	)

	if err != nil {
		opts.Logger.WithError(err).Fatal("Failed to connect to database")
		panic(err)
	}

	if err := db.AutoMigrate(models...); err != nil {
		opts.Logger.WithError(err).Fatal("Failed to migrate database")
		panic(err)
	}

	return db
}
