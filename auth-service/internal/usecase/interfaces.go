package usecase

import (
	"auth-service/internal/domain/user"
	"auth-service/internal/service/jwt"
	"github.com/google/uuid"
)

type StringHasher interface {
	Hash(str string) (string, error)
	Verify(str string, hash string) (bool, error)
}

type AccessTokenIssuer interface {
	IssueAccessToken(id uuid.UUID, username string, role user.UserRole) (string, error)
}

type AccessTokenValidator interface {
	ValidateAccessToken(tokenString string) (*jwt.MyJwtClaims, error)
	ValidateAccessTokenWithRole(tokenString string, role user.UserRole) (bool, error)
}

type RefreshTokenSaver interface {
	GenerateNewAndSave(userId uuid.UUID) error
}
