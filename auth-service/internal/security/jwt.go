package security

import (
	"github.com/golang-jwt/jwt/v5"
	"time"
)

var secret = []byte("topSecret")

const issuer = "auth-service"

func NewTokenIssuer() *BasicTokenIssuer {
	return &BasicTokenIssuer{}
}

type BasicTokenIssuer struct {
}

type MyClaims struct {
	Username string `json:"username"`
	*jwt.RegisteredClaims
}

func (ti *BasicTokenIssuer) CreateAccessToken(username string) (string, error) {
	registeredClaims := jwt.RegisteredClaims{
		Issuer:    issuer,
		IssuedAt:  jwt.NewNumericDate(time.Now()),
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Minute * 60)), // TODO: clean
	}

	claims := &MyClaims{
		Username:         username,
		RegisteredClaims: &registeredClaims,
	}

	var token *jwt.Token = jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	ss, err := token.SignedString(secret)

	if err != nil {
		return "", err
	}

	return ss, nil
}
