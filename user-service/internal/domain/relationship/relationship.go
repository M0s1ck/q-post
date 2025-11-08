package relationship

import (
	"time"

	"github.com/google/uuid"
)

type Relationship struct {
	FollowerId uuid.UUID
	FolloweeId uuid.UUID
	AreFriends bool
	CreatedAt  time.Time
}

func newRelationship(followerId uuid.UUID, followeeId uuid.UUID, areFriends bool, createdAt time.Time) *Relationship {
	return &Relationship{
		FollowerId: followerId,
		FolloweeId: followeeId,
		AreFriends: areFriends,
		CreatedAt:  createdAt}
}
