package security

import (
	"github.com/golang-jwt/jwt/v5"

	"auth-service/internal/domain"
)

type MyJwtClaims struct {
	Username string `json:"username"`
	Role     domain.UserRole
	*jwt.RegisteredClaims
}
