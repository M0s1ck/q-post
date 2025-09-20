package dto

import "github.com/google/uuid"

type UsernamePass struct {
	Username string
	Password string
}

type UserIdAndTokens struct {
	UserId      uuid.UUID
	AccessToken string
}
