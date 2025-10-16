package jwt

import (
	"auth-service/internal/domain/user"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

const (
	authServiceBeingIssuer = "auth-service"
	expireTime             = time.Minute * 60
)

type JwtIssuer struct {
	secret     []byte
	signMethod jwt.SigningMethod
}

func NewJwtIssuer(secret string, signMethod jwt.SigningMethod) *JwtIssuer {
	return &JwtIssuer{
		secret:     []byte(secret),
		signMethod: signMethod,
	}
}

func (ti *JwtIssuer) IssueAccessToken(id uuid.UUID, username string, role user.UserRole) (string, error) {
	registeredClaims := jwt.RegisteredClaims{
		Subject:   id.String(),
		Issuer:    authServiceBeingIssuer,
		IssuedAt:  jwt.NewNumericDate(time.Now()),
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(expireTime)),
	}

	claims := &MyJwtClaims{
		Username:         username,
		Role:             role,
		RegisteredClaims: &registeredClaims,
	}

	var token *jwt.Token = jwt.NewWithClaims(ti.signMethod, claims)

	signedToken, err := token.SignedString(ti.secret)

	if err != nil {
		return "", err
	}

	return signedToken, nil
}
