package jwt

import (
	"github.com/golang-jwt/jwt/v5"

	"auth-service/internal/domain/user"
)

type MyJwtClaims struct {
	Username string `json:"username"`
	Role     user.UserRole
	*jwt.RegisteredClaims
}
