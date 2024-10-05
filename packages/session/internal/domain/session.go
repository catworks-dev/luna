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
	Create(ctx context.Context, session *Session) error

	Get(ctx context.Context, id string) (*Session, error)
	GetByToken(ctx context.Context, token string) (*Session, error)
	List(ctx context.Context) ([]*Session, error)

	Update(ctx context.Context, session *Session) error

	Delete(ctx context.Context, id string) error
	DeleteByToken(ctx context.Context, token string) error
}

type SessionUseCase interface {
	Create(ctx context.Context, rq *CreateSessionRq) (*Session, error)

	Get(ctx context.Context, id string) (*Session, error)
	GetByToken(ctx context.Context, token string) (*Session, error)
	List(ctx context.Context) ([]*Session, error)

	Rename(ctx context.Context, rq *RenameSessionRq) error

	Delete(ctx context.Context, id string) error
}

type CreateSessionRq struct {
	Type DeviceType
	Name string
}

type RenameSessionRq struct {
	Id   string
	Name string
}
