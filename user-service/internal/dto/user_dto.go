package dto

import (
	"github.com/google/uuid"
	"time"
)

type UserDto struct {
	Id           uuid.UUID
	Username     string
	PostKarma    int
	CommentKarma int
	CreatedAt    time.Time
}
