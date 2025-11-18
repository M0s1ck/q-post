package relationships

import (
	"context"

	"github.com/google/uuid"

	"user-service/internal/domain/relationship"
	"user-service/internal/domain/user"
)

type relationshipsGetter interface {
	GetRelationship(userId1 uuid.UUID, userId2 uuid.UUID) (*relationship.Relationship, error)
	GetFriendIds(userId uuid.UUID, offset int, limit int) ([]uuid.UUID, error)
	GetFollowerIds(userId uuid.UUID, offset int, limit int) ([]uuid.UUID, error)
	GetFolloweeIds(userId uuid.UUID, offset int, limit int) ([]uuid.UUID, error)
}

// All write operations require context to support transactions
type relationReaderWriter interface {
	GetRelationship(userId1 uuid.UUID, userId2 uuid.UUID) (*relationship.Relationship, error)
	Add(relation *relationship.Relationship, ctx context.Context) error
	Update(relation *relationship.Relationship, ctx context.Context) error
	Remove(relation *relationship.Relationship, ctx context.Context) error
}

type userGetter interface {
	GetById(id uuid.UUID) (*user.User, error)
	GetUsers([]uuid.UUID) ([]user.User, error)
	ExistsBYId(id uuid.UUID) (bool, error)
}

type userFollowsUpdater interface {
	GetById(id uuid.UUID) (*user.User, error)
	SaveFollowCounts(user *user.User, ctx context.Context) error
}
