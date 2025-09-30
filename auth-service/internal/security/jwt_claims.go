package security

import (
	"auth-service/internal/domain"
	"github.com/golang-jwt/jwt/v5"
)

type MyJwtClaims struct {
	Username string `json:"username"`
	Role     domain.UserRole
	*jwt.RegisteredClaims
}
