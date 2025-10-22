package dto

import (
	"time"

	"github.com/google/uuid"
)

type UserResponse struct {
	Id           uuid.UUID
	Username     string
	PostKarma    int
	CommentKarma int
	Name         *string
	Description  *string
	Birthday     *time.Time
	CreatedAt    time.Time
}

type UserToCreate struct {
	UserId   uuid.UUID `json:"userId"`
	Username string    `json:"username"`
}

type UserDetailsToUpdate struct {
	Name        *string
	Description *string
	Birthday    *time.Time
}
