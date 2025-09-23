package security

import (
	"auth-service/internal/domain"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"time"
)

var secret = []byte("topSecret") // TODO: move to .env

const authServiceBeingIssuer = "auth-service"

func NewTokenIssuer() *BasicTokenIssuer {
	return &BasicTokenIssuer{}
}

type BasicTokenIssuer struct {
}

type MyClaims struct {
	Username string `json:"username"`
	Role     domain.UserRole
	*jwt.RegisteredClaims
}

func (ti *BasicTokenIssuer) CreateAccessToken(id uuid.UUID, username string, role domain.UserRole) (string, error) {
	registeredClaims := jwt.RegisteredClaims{
		Subject:   id.String(),
		Issuer:    authServiceBeingIssuer,
		IssuedAt:  jwt.NewNumericDate(time.Now()),
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Minute * 60)), // TODO: clean
	}

	claims := &MyClaims{
		Username:         username,
		Role:             role,
		RegisteredClaims: &registeredClaims,
	}

	var token *jwt.Token = jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	ss, err := token.SignedString(secret)

	if err != nil {
		return "", err
	}

	return ss, nil
}
