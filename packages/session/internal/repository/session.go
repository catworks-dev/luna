package repository

import (
	"catworks/luna/session/internal/domain"
	"time"
)

type Session struct {
	Id        string            `gorm:"primaryKey"`
	Name      string            `gorm:"not null" gorm:"unique"`
	Type      domain.DeviceType `gorm:"not null" gorm:"type:enum('MOBILE', 'TV')"`
	Token     string            `gorm:"not null" gorm:"uniqueIndex"`
	ExpiresAt time.Time         `gorm:"not null"`
}
