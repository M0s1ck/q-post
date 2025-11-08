package relationships

import (
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

type relationReaderWriter interface {
	GetRelationship(userId1 uuid.UUID, userId2 uuid.UUID) (*relationship.Relationship, error)
	Add(relation *relationship.Relationship) error
	Update(relation *relationship.Relationship) error
	Remove(userId1 uuid.UUID, userId2 uuid.UUID) error
}

type userGetter interface {
	GetById(id uuid.UUID) (*user.User, error)
	GetUsers([]uuid.UUID) ([]user.User, error)
	ExistsBYId(id uuid.UUID) (bool, error)
}

type userFollowsUpdater interface {
	GetById(id uuid.UUID) (*user.User, error)
	SaveFollowCounts(user *user.User) error
}
