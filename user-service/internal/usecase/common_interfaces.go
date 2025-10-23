package usecase

import "github.com/google/uuid"

type AccessTokenValidator interface {
	ValidateApiTokenIssuedAt(token string, issuer string) error
	ValidateUserTokenBySubId(jwt string, userId uuid.UUID) error
	ValidateUserTokenAndGetId(jwt string) (uuid.UUID, error)
}
