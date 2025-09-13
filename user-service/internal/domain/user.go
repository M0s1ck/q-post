package domain

import (
	"github.com/google/uuid"
	"time"
)

type User struct {
	Id           uuid.UUID
	Username     string
	PostKarma    int
	CommentKarma int
	Name         *string // can be null
	Description  *string
	Birthday     *time.Time
	CreatedAt    time.Time
}
