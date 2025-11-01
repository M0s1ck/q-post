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
