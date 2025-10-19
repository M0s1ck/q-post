package refresh

import (
	"github.com/google/uuid"

	"auth-service/internal/domain/user"
)

type RefreshTokenVerifier interface {
	Verify(token uuid.UUID) (userId uuid.UUID, err error)
}

type UserGetter interface {
	GetById(id uuid.UUID) (*user.AuthUser, error)
}
