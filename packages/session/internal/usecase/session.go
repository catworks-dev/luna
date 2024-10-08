package usecase

import (
	"catworks/luna/session/internal/config"
	"catworks/luna/session/internal/domain"
	"context"
	"github.com/google/uuid"
	"time"
)

type sessionUseCaseImpl struct {
	cfg *config.Config

	sessionStorage domain.SessionStorage
	jwtService     domain.JWTService
}

type SessionUseCaseOptions struct {
	Config         *config.Config
	SessionStorage domain.SessionStorage
	JWTService     domain.JWTService
}

func NewSessionUseCase(opts *SessionUseCaseOptions) domain.SessionUseCase {
	return &sessionUseCaseImpl{
		cfg:            opts.Config,
		sessionStorage: opts.SessionStorage,
		jwtService:     opts.JWTService,
	}
}

func (s *sessionUseCaseImpl) Create(ctx context.Context, rq *domain.CreateSessionRq) (*domain.Session, error) {
	id := uuid.New().String()
	token, err := s.jwtService.Generate(ctx, id)
	if err != nil {
		return nil, err
	}

	session := &domain.Session{
		Id:        id,
		Name:      rq.Name,
		Type:      rq.Type,
		Token:     token,
		ExpiresAt: time.Now().Add(s.cfg.SessionTTL),
	}

	err = s.sessionStorage.Create(ctx, session)
	if err != nil {
		return nil, err
	}

	return session, nil
}

func (s *sessionUseCaseImpl) Get(ctx context.Context, id string) (*domain.Session, error) {
	session, err := s.sessionStorage.Get(ctx, id)
	if err != nil {
		return nil, err
	}

	return session, nil
}

func (s *sessionUseCaseImpl) GetByToken(ctx context.Context, token string) (*domain.Session, error) {
	session, err := s.sessionStorage.GetByToken(ctx, token)
	if err != nil {
		return nil, err
	}

	return session, nil
}

func (s *sessionUseCaseImpl) List(ctx context.Context) ([]*domain.Session, error) {
	sessions, err := s.sessionStorage.List(ctx)
	if err != nil {
		return nil, err
	}

	return sessions, nil
}

func (s *sessionUseCaseImpl) Rename(ctx context.Context, rq *domain.RenameSessionRq) error {
	session, err := s.sessionStorage.Get(ctx, rq.Id)
	if err != nil {
		return err
	}

	session.Name = rq.Name
	err = s.sessionStorage.Update(ctx, session)
	if err != nil {
		return err
	}

	return nil
}

func (s *sessionUseCaseImpl) Delete(ctx context.Context, id string) error {
	return s.sessionStorage.Delete(ctx, id)
}
