package usecase

import (
	"github.com/google/uuid"

	"auth-service/internal/domain/user"
	"auth-service/internal/service/jwt"
)

type AccessTokenIssuer interface {
	// IssueAccessToken Issues an access token (jwt) with listed claims
	IssueAccessToken(id uuid.UUID, username string, role user.UserRole) (string, error)
}

type AccessTokenValidator interface {
	ValidateAccessToken(tokenString string) (*jwt.UserJwtClaims, error)
	ValidateAccessTokenWithRole(tokenString string, role user.UserRole) (bool, error)
}
