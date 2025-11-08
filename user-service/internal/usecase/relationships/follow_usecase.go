package relationships

import (
	"errors"
	"github.com/google/uuid"
	"user-service/internal/domain/user"

	"user-service/internal/domain"
	"user-service/internal/domain/relationship"
	"user-service/internal/usecase"
)

type FollowUseCase struct {
	relRepo  relationReaderWriter
	userRepo userFollowsUpdater
	factory  relationship.Factory
	tokenVal usecase.AccessTokenValidator
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
		return u.addFollowership(follower, followee)
	}

	// if followee wants to follow follower -> they become friends
	if followerId == relation.FolloweeId && !relation.AreFriends {
		return u.improveRelationToFriendship(follower, followee, relation)
	}

	return nil
}

func (u *FollowUseCase) addFollowership(follower *user.User, followee *user.User) error {
	relation := u.factory.NewFollowerShip(follower.Id, followee.Id)
	addErr := u.relRepo.Add(relation)
	if addErr != nil {
		return addErr
	}

	follower.FolloweesCount++
	followee.FollowersCount++
	return u.saveGuys(follower, followee)
}

func (u *FollowUseCase) improveRelationToFriendship(follower *user.User, followee *user.User, relation *relationship.Relationship) error {
	relation.AreFriends = true
	updErr := u.relRepo.Update(relation)
	if updErr != nil {
		return updErr
	}

	follower.FollowersCount--
	follower.FriendsCount++
	followee.FolloweesCount--
	followee.FriendsCount++
	return u.saveGuys(follower, followee)
}

func (u *FollowUseCase) saveGuys(follower *user.User, followee *user.User) error {
	saveErr := u.userRepo.SaveFollowCounts(follower)
	if saveErr != nil {
		return saveErr
	}

	saveErr = u.userRepo.SaveFollowCounts(followee)
	return saveErr
}

func NewFollowUseCase(relRepo relationReaderWriter,
	userRepo userFollowsUpdater,
	factory relationship.Factory,
	tokenVal usecase.AccessTokenValidator) *FollowUseCase {
	return &FollowUseCase{
		relRepo:  relRepo,
		userRepo: userRepo,
		factory:  factory,
		tokenVal: tokenVal,
	}
}
