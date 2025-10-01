package security

import (
	"fmt"
	"reflect"

	"github.com/golang-jwt/jwt/v5"

	"auth-service/internal/domain"
)

type JwtValidator struct {
	secret     []byte
	signMethod jwt.SigningMethod
}

func NewJwtValidator(secret string, signMethod jwt.SigningMethod) *JwtValidator {
	return &JwtValidator{
		secret:     []byte(secret),
		signMethod: signMethod,
	}
}

func (v *JwtValidator) ValidateAccessToken(tokenString string) (*MyJwtClaims, error) {
	var claims = MyJwtClaims{}
	token, err := jwt.ParseWithClaims(tokenString, &claims, v.keyFuncSymmetrical)

	if err != nil {
		return nil, err
	}

	if !token.Valid {
		return nil, fmt.Errorf("invalid token")
	}

	return &claims, nil
}

func (v *JwtValidator) ValidateAccessTokenWithRole(tokenString string, role domain.UserRole) (bool, error) {
	claims, err := v.ValidateAccessToken(tokenString)

	if err != nil {
		return false, err
	}

	if claims.Role == role {
		return false, nil
	}

	return true, nil
}

func (v *JwtValidator) keyFuncSymmetrical(token *jwt.Token) (any, error) {
	if reflect.TypeOf(token.Method) == reflect.TypeOf(v.signMethod) {
		return v.secret, nil
	}

	return nil, fmt.Errorf("unexpected signing method: %v", token.Method.Alg())
}
