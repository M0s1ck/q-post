package repository

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"gorm.io/gorm"

	"user-service/internal/domain"
	"user-service/internal/domain/relationship"
)

type FriendRepo struct {
	db *gorm.DB
}

func (f *FriendRepo) GetFriendIds(userId uuid.UUID, offset int, limit int) ([]uuid.UUID, error) {
	ctx := context.Background()

	friendships, err := gorm.G[relationship.Relationship](f.db).
		Where("(follower_id = ? OR followee_id = ?) AND are_friends = ?", userId, userId, true).
		Order("created_at").
		Offset(offset).
		Limit(limit).
		Find(ctx)

	if err != nil {
		return nil, fmt.Errorf("%w: get friends: %v", domain.UnhandledDbError, err)
	}

	var fIds = make([]uuid.UUID, len(friendships))

	for i, fShip := range friendships {
		if fShip.FollowerId == userId {
			fIds[i] = fShip.FolloweeId
		} else {
			fIds[i] = fShip.FollowerId
		}
	}

	return fIds, nil
}

func (f *FriendRepo) GetFollowerIds(userId uuid.UUID, offset int, limit int) ([]uuid.UUID, error) {
	ctx := context.Background()
	var followerIds []uuid.UUID

	err := f.db.WithContext(ctx).
		Model(&relationship.Relationship{}).
		Where("followee_id = ? AND are_friends = ?", userId, false).
		Order("created_at").
		Offset(offset).
		Limit(limit).
		Pluck("follower_id", &followerIds).Error

	if err != nil {
		return nil, fmt.Errorf("%w: get followers: %v", domain.UnhandledDbError, err)
	}

	return followerIds, nil
}

func NewFriendRepo(db *gorm.DB) *FriendRepo {
	return &FriendRepo{db: db}
}
