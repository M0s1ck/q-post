package jwt

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"

	"auth-service/internal/domain/user"
)

const (
	authServiceBeingIssuer = "auth-service"
	expireTime             = time.Minute * 60
)

type JwtIssuer struct {
	userJWTSecret    []byte
	serviceJWTSecret []byte
	signMethod       jwt.SigningMethod
}

func NewJwtIssuer(userJWTSecret string, serviceJWTSecret string, signMethod jwt.SigningMethod) *JwtIssuer {
	return &JwtIssuer{
		userJWTSecret:    []byte(userJWTSecret),
		serviceJWTSecret: []byte(serviceJWTSecret),
		signMethod:       signMethod,
	}
}

func (ti *JwtIssuer) IssueAccessToken(id uuid.UUID, username string, role user.UserRole) (string, error) {
	registeredClaims := jwt.RegisteredClaims{
		Subject:   id.String(),
		Issuer:    authServiceBeingIssuer,
		IssuedAt:  jwt.NewNumericDate(time.Now()),
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(expireTime)),
	}

	claims := &UserJwtClaims{
		Username:         username,
		Role:             role,
		RegisteredClaims: &registeredClaims,
	}

	var token *jwt.Token = jwt.NewWithClaims(ti.signMethod, claims)

	signedToken, err := token.SignedString(ti.userJWTSecret)

	if err != nil {
		return "", err
	}

	return signedToken, nil
}

func (ti *JwtIssuer) IssueJwtForApiService() (string, error) {
	registeredClaims := jwt.RegisteredClaims{
		Subject:   authServiceBeingIssuer,
		Issuer:    authServiceBeingIssuer,
		IssuedAt:  jwt.NewNumericDate(time.Now()),
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(expireTime)),
	}

	claims := &ApiServiceJwtClaims{
		RegisteredClaims: &registeredClaims,
	}

	var token *jwt.Token = jwt.NewWithClaims(ti.signMethod, claims)

	signedToken, err := token.SignedString(ti.serviceJWTSecret)

	if err != nil {
		return "", err
	}

	return signedToken, nil
}
