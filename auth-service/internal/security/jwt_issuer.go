package security

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"

	"auth-service/internal/domain"
)

var secret = []byte("topSecret") // TODO: move to .env

var (
	myJwtSigningMethod = jwt.SigningMethodHS256
)

const authServiceBeingIssuer = "auth-service"

func NewJwtIssuer() *JwtIssuer {
	return &JwtIssuer{}
}

type JwtIssuer struct {
}

func (ti *JwtIssuer) IssueAccessToken(id uuid.UUID, username string, role domain.UserRole) (string, error) {
	registeredClaims := jwt.RegisteredClaims{
		Subject:   id.String(),
		Issuer:    authServiceBeingIssuer,
		IssuedAt:  jwt.NewNumericDate(time.Now()),
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Minute * 60)), // TODO: clean
	}

	claims := &MyJwtClaims{
		Username:         username,
		Role:             role,
		RegisteredClaims: &registeredClaims,
	}

	var token *jwt.Token = jwt.NewWithClaims(myJwtSigningMethod, claims)

	signedToken, err := token.SignedString(secret)

	if err != nil {
		return "", err
	}

	return signedToken, nil
}
