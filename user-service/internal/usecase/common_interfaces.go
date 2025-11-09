package usecase

import (
	"context"

	"github.com/google/uuid"

	"user-service/internal/domain/user"
)

type AccessTokenValidator interface {
	ValidateUserToken(jwt string) error
	ValidateUserTokenAndGetId(token string) (uuid.UUID, error)
	ValidateApiTokenIssuedAt(token string, issuer string) error
	ValidateUserTokenBySubId(token string, userId uuid.UUID) error
	ValidateUserTokenBySubIdOrRole(token string, userId uuid.UUID, role user.GlobalRole) error
}

type UnitOfWork interface {
	// Do doing the writing to bd in fn in one transaction
	Do(ctx context.Context, fn func(ctx context.Context) error) error
}
