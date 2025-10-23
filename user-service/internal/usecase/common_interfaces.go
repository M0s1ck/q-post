package usecase

import (
	"github.com/google/uuid"

	"user-service/internal/domain/user"
)

type AccessTokenValidator interface {
	ValidateUserTokenAndGetId(token string) (uuid.UUID, error)
	ValidateApiTokenIssuedAt(token string, issuer string) error
	ValidateUserTokenBySubId(token string, userId uuid.UUID) error
	ValidateUserTokenBySubIdOrRole(token string, userId uuid.UUID, role user.GlobalRole) error
}
