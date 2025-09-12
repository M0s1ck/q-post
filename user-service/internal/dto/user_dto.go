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
	CreatedAt    time.Time
}

type UserToCreate struct {
	Username string
}
