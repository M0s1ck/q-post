package repository

import (
	"context"
	"errors"
	"fmt"

	"github.com/google/uuid"
	"gorm.io/gorm"

	"user-service/internal/domain"
	"user-service/internal/domain/relationship"
)

type RelationRepo struct {
	*BaseRepo
}

func (r *RelationRepo) GetFriendIds(userId uuid.UUID, offset int, limit int) ([]uuid.UUID, error) {
	ctx := context.Background()

	friendships, err := gorm.G[relationship.Relationship](r.db).
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

func (r *RelationRepo) GetFollowerIds(userId uuid.UUID, offset int, limit int) ([]uuid.UUID, error) {
	ctx := context.Background()
	var followerIds []uuid.UUID

	err := r.db.WithContext(ctx).
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

func (r *RelationRepo) GetFolloweeIds(userId uuid.UUID, offset int, limit int) ([]uuid.UUID, error) {
	ctx := context.Background()
	var followeeIds []uuid.UUID

	err := r.db.WithContext(ctx).
		Model(&relationship.Relationship{}).
		Where("follower_id = ? AND are_friends = ?", userId, false).
		Order("created_at").
		Offset(offset).
		Limit(limit).
		Pluck("followee_id", &followeeIds).Error

	if err != nil {
		return nil, fmt.Errorf("%w: get followees: %v", domain.UnhandledDbError, err)
	}

	return followeeIds, nil
}

func (r *RelationRepo) GetRelationship(userId1 uuid.UUID, userId2 uuid.UUID) (*relationship.Relationship, error) {
	ctx := context.Background()

	rel, err := gorm.G[relationship.Relationship](r.db).
		Where("(follower_id = ? AND followee_id = ?) OR (followee_id = ? AND follower_id = ?)",
			userId1, userId2, userId1, userId2).
		First(ctx)

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, fmt.Errorf("%w: get relationship: %v", domain.ErrNotFound, err)
	}

	if err != nil {
		return nil, fmt.Errorf("%w: get relationship: %v", domain.UnhandledDbError, err)
	}

	return &rel, nil
}

func (r *RelationRepo) Add(relation *relationship.Relationship, ctx context.Context) error {
	tx := r.getTx(ctx)
	err := gorm.G[relationship.Relationship](tx).Create(ctx, relation)

	if err != nil {
		return fmt.Errorf("%w: add relationship: %v", domain.UnhandledDbError, err)
	}

	return nil
}

func (r *RelationRepo) Update(relation *relationship.Relationship, ctx context.Context) error {
	tx := r.getTx(ctx)
	_, err := gorm.G[relationship.Relationship](tx).
		Where("(follower_id = ? AND followee_id = ?) OR (followee_id = ? AND follower_id = ?)",
			relation.FollowerId, relation.FolloweeId, relation.FollowerId, relation.FolloweeId).
		Select("are_friends").
		Updates(ctx, relationship.Relationship{AreFriends: relation.AreFriends})

	if err != nil {
		return fmt.Errorf("%w: update relationship: %v", domain.UnhandledDbError, err)
	}

	return nil
}

func (r *RelationRepo) Remove(relation *relationship.Relationship, ctx context.Context) error {
	tx := r.getTx(ctx)
	_, err := gorm.G[relationship.Relationship](tx).
		Where("follower_id = ? AND followee_id = ?", relation.FollowerId, relation.FolloweeId).
		Delete(ctx)

	if err != nil {
		return fmt.Errorf("%w: remove relationship: %v", domain.UnhandledDbError, err)
	}

	return nil
}

func (r *RelationRepo) RemoveByIds(userId1 uuid.UUID, userId2 uuid.UUID) error {
	ctx := context.Background()
	_, err := gorm.G[relationship.Relationship](r.db).
		Where("(follower_id = ? AND followee_id = ?) OR (followee_id = ? AND follower_id = ?)",
			userId1, userId2, userId1, userId2).
		Delete(ctx)

	if err != nil {
		return fmt.Errorf("%w: remove relationship: %v", domain.UnhandledDbError, err)
	}

	return nil
}

func NewRelationRepo(baseRepo *BaseRepo) *RelationRepo {
	return &RelationRepo{BaseRepo: baseRepo}
}
