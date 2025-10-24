package jwt

import (
	"github.com/golang-jwt/jwt/v5"

	"user-service/internal/domain/user"
)

type ApiServiceClaims struct {
	*jwt.RegisteredClaims
}

type UserClaims struct {
	Username string `json:"username"`
	Role     user.GlobalRole
	*jwt.RegisteredClaims
}
