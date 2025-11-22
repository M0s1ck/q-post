package auth

import (
	"github.com/google/uuid"
)

type RefreshTokenGenerator interface {
	// GenerateNewAndSave Generates uuid refresh token for user and saves it to db
	GenerateNewAndSave(userId uuid.UUID) (uuid.UUID, error)
}
