package dto

import "github.com/google/uuid"

type UserToCreate struct {
	UserId   uuid.UUID `json:"userId"`
	Username string    `json:"username"`
}
