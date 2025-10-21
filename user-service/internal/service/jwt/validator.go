package jwt

import (
	"fmt"
	"reflect"

	"github.com/golang-jwt/jwt/v5"
)

type Validator struct {
	userSecret   []byte
	apiSecretKey []byte
	signMethod   jwt.SigningMethod
}

func NewValidator(secret string, apiSecretKey string, signMethod jwt.SigningMethod) *Validator {
	return &Validator{
		userSecret:   []byte(secret),
		apiSecretKey: []byte(apiSecretKey),
		signMethod:   signMethod,
	}
}

func (v *Validator) ValidateApiToken(tokenStr string) (*ApiServiceClaims, error) {
	var claims = ApiServiceClaims{}
	token, err := jwt.ParseWithClaims(tokenStr, &claims, v.keyFuncApiSymmetrical)

	if err != nil {
		return nil, err
	}

	if !token.Valid {
		return nil, fmt.Errorf("invalid token")
	}

	return &claims, nil
}

func (v *Validator) ValidateTokenIssuedAt(jwt string, issuer string) error {
	claims, err := v.ValidateApiToken(jwt)

	if err != nil {
		return err
	}

	if claims.Issuer != issuer {
		return fmt.Errorf("user create: unexpected jwt issuer: %v", claims.Issuer)
	}

	return nil
}

func (v *Validator) keyFuncApiSymmetrical(token *jwt.Token) (any, error) {
	if reflect.TypeOf(token.Method) == reflect.TypeOf(v.signMethod) {
		return v.apiSecretKey, nil
	}

	return nil, fmt.Errorf("unexpected signing method: %v", token.Method.Alg())
}
