package relationships

import (
	"errors"
	"fmt"

	"github.com/google/uuid"

	"user-service/internal/domain"
	"user-service/internal/domain/relationship"
	"user-service/internal/usecase"
)

type FollowUseCase struct {
	relRepo  relationReaderWriter
	userRepo userGetter
	factory  relationship.Factory
	tokenVal usecase.AccessTokenValidator
}

func (u *FollowUseCase) Follow(followeeId uuid.UUID, token string) error {
	followerId, tokenErr := u.tokenVal.ValidateUserTokenAndGetId(token)
	if tokenErr != nil {
		return tokenErr
	}

	exists, usErr := u.userRepo.ExistsBYId(followeeId)
	if usErr != nil {
		return usErr
	}
	if !exists {
		return fmt.Errorf("followee not found: %w", domain.ErrNotFound)
	}

	relation, relErr := u.relRepo.GetRelationship(followerId, followeeId)
	if relErr != nil && !errors.Is(relErr, domain.ErrNotFound) {
		return relErr
	}

	// if no relationship, we add it
	if errors.Is(relErr, domain.ErrNotFound) {
		relation = u.factory.NewFollowerShip(followerId, followeeId)
		addErr := u.relRepo.Add(relation)
		if addErr != nil {
			return addErr
		}
		return nil
	}

	// if followee wants to follow follower -> they become friends
	if followerId == relation.FolloweeId && !relation.AreFriends {
		relation.AreFriends = true
		updErr := u.relRepo.Update(relation)
		if updErr != nil {
			return updErr
		}
		return nil
	}

	return nil
}

func NewFollowUseCase(relRepo relationReaderWriter,
	userRepo userGetter,
	factory relationship.Factory,
	tokenVal usecase.AccessTokenValidator) *FollowUseCase {
	return &FollowUseCase{
		relRepo:  relRepo,
		userRepo: userRepo,
		factory:  factory,
		tokenVal: tokenVal,
	}
}
