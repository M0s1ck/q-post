package repository

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"gorm.io/gorm"

	"user-service/internal/domain"
	"user-service/internal/domain/friendship"
)

type FriendRepo struct {
	db *gorm.DB
}

func (f FriendRepo) GetFriends(userId uuid.UUID, offset int, limit int) ([]uuid.UUID, error) {
	ctx := context.Background()

	friendships, err := gorm.G[friendship.Friend](f.db).
		Where("(follower_id = ? OR followee_id = ?) AND are_friends = ?", userId, userId, true).
		Order("created_at").
		Offset(offset).
		Limit(limit).
		Find(ctx)

	var fIds = make([]uuid.UUID, len(friendships))

	for i, fShip := range friendships {
		if fShip.FollowerId == userId {
			fIds[i] = fShip.FolloweeId
		} else {
			fIds[i] = fShip.FollowerId
		}
	}

	if err != nil {
		return nil, fmt.Errorf("%w: get friends: %v", domain.UnhandledDbError, err)
	}

	return fIds, nil
}

func NewFriendRepo(db *gorm.DB) *FriendRepo {
	return &FriendRepo{db: db}
}
