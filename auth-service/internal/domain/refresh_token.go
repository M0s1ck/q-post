package domain

import (
	"github.com/google/uuid"
	"time"
)

// Many to one with users cause they can use app from different devices

type RefreshToken struct {
	id        uuid.UUID
	tokenHash string
	userId    uuid.UUID
	expiresAt time.Time
}
