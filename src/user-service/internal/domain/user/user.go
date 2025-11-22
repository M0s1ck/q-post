package user

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	Id             uuid.UUID  `gorm:"column:id;primaryKey"`
	Username       string     `gorm:"column:username"`
	Name           *string    `gorm:"column:name"`
	Description    *string    `gorm:"column:description"`
	Birthday       *time.Time `gorm:"column:birthday"`
	FriendsCount   int        `gorm:"column:friends_count"`
	FollowersCount int        `gorm:"column:followers_count"`
	FolloweesCount int        `gorm:"column:followees_count"`
	PostKarma      int        `gorm:"column:post_karma"`
	CommentKarma   int        `gorm:"column:comment_karma"`
	CreatedAt      time.Time  `gorm:"column:created_at"`
}
