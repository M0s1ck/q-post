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
		return tokenErr // TODO: change these errors to ErrToken to avoid returning 500
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
	if followerId == relation.FolloweeId && !relation.AreFriends { // TODO: test
		ctx := context.Background()
		return u.unitOfWork.Do(ctx, func(txCtx context.Context) error {
			return u.improveRelationToFriendship(follower, followee, relation, txCtx)
		})
	}

	return nil
}

func (u *FollowUseCase) addFollowership(follower *user.User, followee *user.User, ctx context.Context) error {
	relation := u.factory.NewFollowerShip(follower.Id, followee.Id)
	addErr := u.relRepo.Add(ctx, relation)
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
	updErr := u.relRepo.Update(relation)
	if updErr != nil {
		return updErr
	}

	follower.FollowersCount--
	follower.FriendsCount++
	followee.FolloweesCount--
	followee.FriendsCount++
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
