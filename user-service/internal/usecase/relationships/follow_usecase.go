package relationships

import (
	"context"
	"errors"

	"github.com/google/uuid"

	"user-service/internal/domain"
	"user-service/internal/domain/relationship"
	"user-service/internal/domain/user"
	"user-service/internal/usecase"
)

type FollowUseCase struct {
	relRepo    relationReaderWriter
	userRepo   userFollowsUpdater
	factory    relationship.Factory
	unitOfWork usecase.UnitOfWork
	tokenVal   usecase.AccessTokenValidator
}

func (u *FollowUseCase) Follow(followeeId uuid.UUID, token string) error {
	followerId, tokenErr := u.tokenVal.ValidateUserTokenAndGetId(token)
	if tokenErr != nil {
		return tokenErr
	}

	if followeeId == followerId {
		return domain.ErrSelfFollow
	}

	followee, usErr := u.userRepo.GetById(followeeId)
	if usErr != nil {
		return usErr
	}

	follower, usErr := u.userRepo.GetById(followerId)
	if usErr != nil {
		return usErr
	}

	relation, relErr := u.relRepo.GetRelationship(followerId, followeeId)
	if relErr != nil && !errors.Is(relErr, domain.ErrNotFound) {
		return relErr
	}

	// if no relationship, we add it
	if errors.Is(relErr, domain.ErrNotFound) {
		ctx := context.Background()
		return u.unitOfWork.Do(ctx, func(txCtx context.Context) error {
			return u.addFollowership(follower, followee, txCtx)
		})
	}

	// if followee wants to follow follower -> they become friends
	if followerId == relation.FolloweeId && !relation.AreFriends {
		ctx := context.Background()
		return u.unitOfWork.Do(ctx, func(txCtx context.Context) error {
			return u.improveRelationToFriendship(follower, followee, relation, txCtx)
		})
	}

	return nil
}

func (u *FollowUseCase) Unfollow(followeeId uuid.UUID, token string) error {
	followerId, tokenErr := u.tokenVal.ValidateUserTokenAndGetId(token)
	if tokenErr != nil {
		return tokenErr
	}

	if followeeId == followerId {
		return domain.ErrSelfFollow
	}

	followee, usErr := u.userRepo.GetById(followeeId)
	if usErr != nil {
		return usErr
	}

	follower, usErr := u.userRepo.GetById(followerId)
	if usErr != nil {
		return usErr
	}

	relation, relErr := u.relRepo.GetRelationship(followerId, followeeId)
	if relErr != nil {
		return relErr
	}

	// if were friends -> they're nobody to each other now
	if relation.AreFriends {
		ctx := context.Background()
		return u.unitOfWork.Do(ctx, func(txCtx context.Context) error {
			return u.removeFriendship(follower, followee, relation, txCtx)
		})
	}

	// otherwise guy was following -> unfollows
	if relation.FollowerId == followerId && relation.FolloweeId == followeeId {
		ctx := context.Background()
		return u.unitOfWork.Do(ctx, func(txCtx context.Context) error {
			return u.removeFollowership(follower, followee, relation, txCtx)
		})
	}

	return nil
}

func (u *FollowUseCase) addFollowership(follower *user.User, followee *user.User, ctx context.Context) error {
	relation := u.factory.NewFollowerShip(follower.Id, followee.Id)
	addErr := u.relRepo.Add(relation, ctx)
	if addErr != nil {
		return addErr
	}

	follower.FolloweesCount++
	followee.FollowersCount++
	return u.saveGuys(follower, followee, ctx)
}

func (u *FollowUseCase) improveRelationToFriendship(follower *user.User, followee *user.User,
	relation *relationship.Relationship, ctx context.Context) error {
	relation.AreFriends = true
	updErr := u.relRepo.Update(relation, ctx)
	if updErr != nil {
		return updErr
	}

	follower.FollowersCount--
	follower.FriendsCount++
	followee.FolloweesCount--
	followee.FriendsCount++
	return u.saveGuys(follower, followee, ctx)
}

func (u *FollowUseCase) removeFollowership(follower *user.User, followee *user.User,
	relation *relationship.Relationship, ctx context.Context) error {
	relErr := u.relRepo.Remove(relation, ctx)
	if relErr != nil {
		return relErr
	}

	follower.FolloweesCount--
	followee.FollowersCount--
	return u.saveGuys(follower, followee, ctx)
}

func (u *FollowUseCase) removeFriendship(follower *user.User, followee *user.User,
	relation *relationship.Relationship, ctx context.Context) error {
	relErr := u.relRepo.Remove(relation, ctx)
	if relErr != nil {
		return relErr
	}

	follower.FriendsCount--
	followee.FriendsCount--
	return u.saveGuys(follower, followee, ctx)
}

func (u *FollowUseCase) saveGuys(follower *user.User, followee *user.User, ctx context.Context) error {
	saveErr := u.userRepo.SaveFollowCounts(follower, ctx)
	if saveErr != nil {
		return saveErr
	}

	saveErr = u.userRepo.SaveFollowCounts(followee, ctx)
	return saveErr
}

func NewFollowUseCase(relRepo relationReaderWriter,
	userRepo userFollowsUpdater,
	factory relationship.Factory,
	uow usecase.UnitOfWork,
	tokenVal usecase.AccessTokenValidator) *FollowUseCase {
	return &FollowUseCase{
		relRepo:    relRepo,
		userRepo:   userRepo,
		factory:    factory,
		unitOfWork: uow,
		tokenVal:   tokenVal,
	}
}
