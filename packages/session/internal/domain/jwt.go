package domain

import "context"

type JWTService interface {
	Generate(ctx context.Context, id string) (string, error)
}
