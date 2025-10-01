package usecase

import (
	"github.com/google/uuid"

	"auth-service/internal/domain"
	"auth-service/internal/security"
)

type PasswordHasher interface {
	Hash(password string) (string, error)
	Verify(password string, hash string) (bool, error)
}

type TokenIssuer interface {
	IssueAccessToken(id uuid.UUID, username string, role domain.UserRole) (string, error)
}

type TokenValidator interface {
	ValidateAccessToken(tokenString string) (*security.MyJwtClaims, error)
	ValidateAccessTokenWithRole(tokenString string, role domain.UserRole) (bool, error)
}
