package user

import (
	"github.com/google/uuid"
)

type AuthUser struct {
	Id             uuid.UUID
	Username       string
	Email          *string
	HashedPassword string
	Role           UserRole `gorm:"type:smallint;not null;default:0"`
}

func NewAuthUser(id uuid.UUID, username string, email *string, passHash string, role UserRole) *AuthUser {
	return &AuthUser{
		Id:             id,
		Username:       username,
		Email:          email,
		HashedPassword: passHash,
		Role:           role,
	}
}
