package dto

import (
	"time"

	"github.com/google/uuid"
)

// UserResponse used in user's profile
type UserResponse struct {
	Id             uuid.UUID  `json:"id" example:"550e8400-e29b-41d4-a716-446655440000"`
	Username       string     `json:"username" example:"M0s1ck"`
	Name           *string    `json:"name" example:"John Doe"`
	Description    *string    `json:"description" example:"I love basketball and films"`
	Birthday       *time.Time `json:"birthday" example:"2006-01-02"`
	FriendsCount   int        `json:"friendsCount" example:"120"`
	FollowersCount int        `json:"followersCount" example:"540"`
	FolloweesCount int        `json:"followeesCount" example:"320"`
	PostKarma      int        `json:"postKarma" example:"1500"`
	CommentKarma   int        `json:"commentKarma" example:"800"`
	CreatedAt      time.Time  `json:"createdAt" example:"2024-11-01T12:34:56Z"`
}

// UserSummary used to show user in lists, comments etc.
type UserSummary struct {
	Id       uuid.UUID `json:"id" example:"550e8400-e29b-41d4-a716-446655440000"`
	Username string    `json:"username" example:"M0s1ck"`
}

type UserToCreate struct {
	UserId   uuid.UUID `json:"userId" example:"550e8400-e29b-41d4-a716-446655440000"`
	Username string    `json:"username" example:"M0s1ck"`
}

// UserDetailStr these can be updated by user
type UserDetailStr struct {
	Name        *string `json:"name" example:"John Doe"`
	Description *string `json:"description" example:"I love ball and films'"`
	Birthday    *string `json:"birthday" example:"2006-01-02"`
}
