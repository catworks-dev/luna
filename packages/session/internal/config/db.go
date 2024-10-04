package config

import (
	"catworks/luna/session/internal/repository"
	"context"
	"github.com/sirupsen/logrus"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"time"
)

var models = []interface{}{
	&repository.Session{},
}

func NewGorm(config *Config, logger *logrus.Logger) *gorm.DB {
	db, err := gorm.Open(
		sqlite.Open(config.StoragePath),
		&gorm.Config{
			Logger: &DbLogger{Logger: logger},
		},
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

type DbLogger struct {
	Logger *logrus.Logger
}

func (d DbLogger) LogMode(level logger.LogLevel) logger.Interface {
	l := logrus.New()

	var lvl logrus.Level
	switch level {
	case logger.Silent:
		lvl = logrus.TraceLevel
	case logger.Error:
		lvl = logrus.ErrorLevel
	case logger.Warn:
		lvl = logrus.WarnLevel
	case logger.Info:
		lvl = logrus.InfoLevel
	}

	l.SetLevel(lvl)
	l.SetFormatter(d.Logger.Formatter)
	l.SetOutput(d.Logger.Out)

	return DbLogger{Logger: l}
}

func (d DbLogger) Info(ctx context.Context, s string, i ...interface{}) {
	d.Logger.WithContext(ctx).WithField("data", i).Info(s)
}

func (d DbLogger) Warn(ctx context.Context, s string, i ...interface{}) {
	d.Logger.WithContext(ctx).WithField("data", i).Warn(s)
}

func (d DbLogger) Error(ctx context.Context, s string, i ...interface{}) {
	d.Logger.WithContext(ctx).WithField("data", i).Error(s)
}

func (d DbLogger) Trace(ctx context.Context, begin time.Time, fc func() (sql string, rowsAffected int64), err error) {
	sql, rowsAffected := fc()
	if err != nil {
		d.Logger.WithContext(ctx).WithError(err).WithField("sql", sql).WithTime(begin).Error("Failed to execute query")
	} else {
		d.Logger.WithContext(ctx).WithField("sql", sql).WithField("rowsAffected", rowsAffected).Info("Query executed")
	}
}
