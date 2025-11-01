package relationships

import (
	"github.com/google/uuid"

	"user-service/internal/domain/user"
	"user-service/internal/dto"
	"user-service/internal/mapper"
	"user-service/internal/usecase"
)

const DefaultPageSize = 20

type relationshipsGetter interface {
	GetFriendIds(userId uuid.UUID, offset int, limit int) ([]uuid.UUID, error)
	GetFollowerIds(userId uuid.UUID, offset int, limit int) ([]uuid.UUID, error)
	GetFolloweeIds(userId uuid.UUID, offset int, limit int) ([]uuid.UUID, error)
}

type usersGetter interface {
	GetUsers([]uuid.UUID) ([]user.User, error)
}

type GetRelationshipsUseCase struct {
	relationshipRepo relationshipsGetter
	userRepo         usersGetter
	tokenVal         usecase.AccessTokenValidator
}

func NewGetRelationshipsUseCase(fRepo relationshipsGetter, uRepo usersGetter, tVal usecase.AccessTokenValidator) *GetRelationshipsUseCase {
	return &GetRelationshipsUseCase{
		relationshipRepo: fRepo,
		userRepo:         uRepo,
		tokenVal:         tVal,
	}
}

func (u *GetRelationshipsUseCase) GetFriends(userId uuid.UUID, page int, pageSize int, token string) ([]dto.UserSummary, error) {
	tokenErr := u.tokenVal.ValidateUserToken(token)
	if tokenErr != nil {
		return nil, tokenErr
	}

	offset := pageSize * page

	friendIds, err := u.relationshipRepo.GetFriendIds(userId, offset, pageSize)
	if err != nil {
		return nil, err
	}

	friends, err := u.userRepo.GetUsers(friendIds)
	if err != nil {
		return nil, err
	}

	summaries := mapper.GetUserSummaries(friends)
	return summaries, nil
}

func (u *GetRelationshipsUseCase) GetFollowers(userId uuid.UUID, page int, pageSize int, token string) ([]dto.UserSummary, error) {
	tokenErr := u.tokenVal.ValidateUserToken(token)
	if tokenErr != nil {
		return nil, tokenErr
	}

	offset := pageSize * page

	followerIds, err := u.relationshipRepo.GetFollowerIds(userId, offset, pageSize)
	if err != nil {
		return nil, err
	}

	followers, err := u.userRepo.GetUsers(followerIds)
	if err != nil {
		return nil, err
	}

	summaries := mapper.GetUserSummaries(followers)
	return summaries, nil
}

func (u *GetRelationshipsUseCase) GetFollowees(userId uuid.UUID, page int, pageSize int, token string) ([]dto.UserSummary, error) {
	tokenErr := u.tokenVal.ValidateUserToken(token)
	if tokenErr != nil {
		return nil, tokenErr
	}

	offset := pageSize * page

	followeeIds, err := u.relationshipRepo.GetFolloweeIds(userId, offset, pageSize)
	if err != nil {
		return nil, err
	}

	followers, err := u.userRepo.GetUsers(followeeIds)
	if err != nil {
		return nil, err
	}

	summaries := mapper.GetUserSummaries(followers)
	return summaries, nil
}
