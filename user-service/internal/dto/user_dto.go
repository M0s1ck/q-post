package dto

import (
	"github.com/google/uuid"
	"time"
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
	Username string
}

type UserDetailsToUpdate struct {
	Name        *string
	Description *string
	Birthday    *time.Time
}
