package dto

import "github.com/google/uuid"

type RefreshDto struct {
	Token uuid.UUID `json:"refresh_token"`
}
