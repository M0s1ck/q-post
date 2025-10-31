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

type UserSummary struct {
	Id       uuid.UUID
	Username string
}

type UserToCreate struct {
	UserId   uuid.UUID `json:"userId"`
	Username string    `json:"username"`
}

type UserDetailStr struct {
	Name        *string `json:"name" example:"John Doe"`
	Description *string `json:"description" example:"I love ball and films'"`
	Birthday    *string `json:"birthday" example:"2006-01-02"`
}
