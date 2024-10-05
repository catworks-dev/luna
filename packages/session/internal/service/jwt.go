package service

import (
	"catworks/luna/session/internal/config"
	"catworks/luna/session/internal/domain"
	"context"
	"github.com/golang-jwt/jwt/v5"
	"time"
)

type jwtServiceImpl struct {
	cfg *config.Config
}

// NewJWTService - конструктор для jwtServiceImpl
func NewJWTService(cfg *config.Config) domain.JWTService {
	return &jwtServiceImpl{cfg: cfg}
}

// Generate - метод для генерации JWT
func (s *jwtServiceImpl) Generate(_ context.Context, id string) (string, error) {
	// Создаем токен
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"exp": time.Now().Add(s.cfg.SessionTTL).Unix(), // срок действия токена
		"iat": time.Now().Unix(),                       // время создания токена
		"sub": id,                                      // идентификатор пользователя
	})

	// Подписываем токен секретным ключом
	tokenString, err := token.SignedString([]byte(s.cfg.JwtSecret))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
