package repository

import (
	"catworks/luna/session/internal/domain"
	"context"
	"gorm.io/gorm"
	"time"
)

type Session struct {
	Id        string            `gorm:"primaryKey"`
	Name      string            `gorm:"not null" gorm:"unique"`
	Type      domain.DeviceType `gorm:"not null" gorm:"type:enum('MOBILE', 'TV')"`
	Token     string            `gorm:"not null" gorm:"uniqueIndex"`
	ExpiresAt time.Time         `gorm:"not null"`
}

func NewSessionRepository(db *gorm.DB) domain.SessionStorage {
	return sessionStorageImpl{db: db}
}

type sessionStorageImpl struct {
	db *gorm.DB
}

func (s sessionStorageImpl) Create(ctx context.Context, session *domain.Session) error {
	return s.db.WithContext(ctx).Create(
		s.sessionFromDomain(session),
	).Error
}

func (s sessionStorageImpl) Get(ctx context.Context, id string) (*domain.Session, error) {
	var session Session
	err := s.db.WithContext(ctx).Where("id = ?", id).First(&session).Error
	if err != nil {
		return nil, err
	}

	return s.sessionToDomain(&session), nil
}

func (s sessionStorageImpl) GetByToken(ctx context.Context, token string) (*domain.Session, error) {
	var session Session
	err := s.db.WithContext(ctx).Where("token = ?", token).First(&session).Error
	if err != nil {
		return nil, err
	}

	return s.sessionToDomain(&session), nil
}

func (s sessionStorageImpl) List(ctx context.Context) ([]*domain.Session, error) {
	var sessions []Session
	err := s.db.WithContext(ctx).Find(&sessions).Error
	if err != nil {
		return nil, err
	}

	var result []*domain.Session
	for _, session := range sessions {
		result = append(result, s.sessionToDomain(&session))
	}

	return result, nil
}

func (s sessionStorageImpl) Update(ctx context.Context, session *domain.Session) error {
	return s.db.WithContext(ctx).Model(&Session{}).Where("id = ?", session.Id).Updates(
		s.sessionFromDomain(session),
	).Error
}

func (s sessionStorageImpl) Delete(ctx context.Context, id string) error {
	return s.db.WithContext(ctx).Delete(&Session{}, id).Error
}

func (s sessionStorageImpl) DeleteByToken(ctx context.Context, token string) error {
	return s.db.WithContext(ctx).Delete(&Session{}, "token = ?", token).Error
}

func (s sessionStorageImpl) sessionFromDomain(session *domain.Session) *Session {
	return &Session{
		Id:        session.Id,
		Name:      session.Name,
		Type:      session.Type,
		Token:     session.Token,
		ExpiresAt: session.ExpiresAt,
	}
}

func (s sessionStorageImpl) sessionToDomain(session *Session) *domain.Session {
	return &domain.Session{
		Id:        session.Id,
		Name:      session.Name,
		Type:      session.Type,
		Token:     session.Token,
		ExpiresAt: session.ExpiresAt,
	}
}
