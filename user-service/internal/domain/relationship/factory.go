package relationship

import (
	"time"

	"github.com/google/uuid"
)

type Factory interface {
	NewFollowerShip(followerId uuid.UUID, followeeId uuid.UUID) *Relationship
}

type DefaultFactory struct {
}

func NewDefaultFactory() *DefaultFactory {
	return &DefaultFactory{}
}

func (f *DefaultFactory) NewFollowerShip(followerId uuid.UUID, followeeId uuid.UUID) *Relationship {
	return &Relationship{
		FollowerId: followerId,
		FolloweeId: followeeId,
		AreFriends: false,
		CreatedAt:  time.Now(),
	}
}
