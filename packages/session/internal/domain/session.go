package domain

import (
	"context"
	"time"
)

type DeviceType int

const (
	MOBILE DeviceType = iota
	TV
)

type Session struct {
	Id        string
	Name      string
	Type      DeviceType
	Token     string
	ExpiresAt time.Time
}

type SessionStorage interface {
	Get(ctx context.Context, id string) (*Session, error)
	GetByToken(ctx context.Context, token string) (*Session, error)
	List(ctx context.Context) ([]*Session, error)

	Create(ctx context.Context, session *Session) error

	Update(ctx context.Context, session *Session) error

	Delete(ctx context.Context, id string) error
	DeleteByToken(ctx context.Context, token string) error
}
