package jwt

import (
	"github.com/golang-jwt/jwt/v5"

	"auth-service/internal/domain/user"
)

type UserJwtClaims struct {
	Username string `json:"username"`
	Role     user.UserRole
	*jwt.RegisteredClaims
}

type ApiServiceJwtClaims struct {
	*jwt.RegisteredClaims
}
