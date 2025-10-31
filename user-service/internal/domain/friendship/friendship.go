package friendship

import (
	"time"

	"github.com/google/uuid"
)

type Friend struct {
	FollowerId uuid.UUID
	FolloweeId uuid.UUID
	AreFriends bool
	CreatedAt  time.Time
}
