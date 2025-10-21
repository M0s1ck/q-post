package domain

import (
	"time"

	"github.com/google/uuid"
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
