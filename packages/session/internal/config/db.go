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

func NewGorm(config *Config, logger *logrus.Logger) *gorm.DB {
	gormConfig := &gorm.Config{}
	if config.logDB {
		gormConfig.Logger = gl.New()
	}
	db, err := gorm.Open(
		sqlite.Open(config.StoragePath),
		gormConfig,
	)

	if err != nil {
		logger.WithError(err).Fatal("Failed to connect to database")
		panic(err)
	}

	if err := db.AutoMigrate(models...); err != nil {
		logger.WithError(err).Fatal("Failed to migrate database")
		panic(err)
	}

	return db
}
