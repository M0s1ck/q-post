package domain

import "github.com/google/uuid"

type AuthUser struct {
	Id             uuid.UUID
	Username       string
	Email          *string
	HashedPassword string
	Role           UserRole `gorm:"type:smallint;not null;default:0"`
}
