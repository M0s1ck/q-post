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
	Id       uuid.UUID
	Username string
}

type UserDetailsToUpdate struct {
	Name        *string
	Description *string
	Birthday    *time.Time
}
