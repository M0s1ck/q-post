package security

import (
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"reflect"
)

type JwtValidator struct {
}

func NewJwtValidator() *JwtValidator {
	return &JwtValidator{}
}

func (v *JwtValidator) ValidateAccessToken(tokenString string) (*MyJwtClaims, error) {
	var claims = MyJwtClaims{}
	token, err := jwt.ParseWithClaims(tokenString, &claims, keyFunc)

	if err != nil {
		return nil, err
	}

	if !token.Valid {
		return nil, fmt.Errorf("invalid token")
	}

	return &claims, nil
}

func keyFunc(token *jwt.Token) (any, error) {
	if reflect.TypeOf(token.Method) == reflect.TypeOf(myJwtSigningMethod) {
		return secret, nil
	}

	return nil, fmt.Errorf("unexpected signing method: %v", token.Method.Alg())
}
