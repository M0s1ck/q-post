package usecase

import (
	"github.com/google/uuid"

	"auth-service/internal/domain"
)

type PasswordHasher interface {
	Hash(password string) (string, error)
	Verify(password string, hash string) (bool, error)
}

type TokenIssuer interface {
	CreateAccessToken(id uuid.UUID, username string, role domain.UserRole) (string, error)
}
