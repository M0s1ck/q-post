package refresh

import (
	"time"

	"github.com/google/uuid"
)

const refreshTokenLifeSpan = time.Hour * 24 * 30

// Many to one with users cause they can use app from different devices

type RefreshToken struct {
	Id        uuid.UUID
	TokenHash string
	UserId    uuid.UUID
	ExpiresAt time.Time
}

func NewRefreshToken(tokenHash string, userId uuid.UUID) RefreshToken {
	return RefreshToken{
		Id:        uuid.New(),
		TokenHash: tokenHash,
		UserId:    userId,
		ExpiresAt: time.Now().Add(refreshTokenLifeSpan),
	}
}
