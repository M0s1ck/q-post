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
	CreatedAt    time.Time
}
