package jwt

import (
	"auth-service/internal/domain/user"
	"github.com/golang-jwt/jwt/v5"
)

type MyJwtClaims struct {
	Username string `json:"username"`
	Role     user.UserRole
	*jwt.RegisteredClaims
}
