package friends

import (
	"github.com/google/uuid"

	"user-service/internal/domain/user"
	"user-service/internal/dto"
	"user-service/internal/mapper"
	"user-service/internal/usecase"
)

const DefaultPageSize = 20

type friendsGetter interface {
	GetFriends(userId uuid.UUID, offset int, limit int) ([]uuid.UUID, error)
}

type usersGetter interface {
	GetUsers([]uuid.UUID) ([]user.User, error)
}

type GetFriendsUseCase struct {
	friendsRepo friendsGetter
	userRepo    usersGetter
	tokenVal    usecase.AccessTokenValidator
}

func NewGetFriendsUseCase(fRepo friendsGetter, uRepo usersGetter, tVal usecase.AccessTokenValidator) *GetFriendsUseCase {
	return &GetFriendsUseCase{
		friendsRepo: fRepo,
		userRepo:    uRepo,
		tokenVal:    tVal,
	}
}

func (u *GetFriendsUseCase) GetFriends(userId uuid.UUID, page int, pageSize int, token string) ([]dto.UserSummary, error) {
	_, tokenErr := u.tokenVal.ValidateUserTokenAndGetId(token)
	if tokenErr != nil {
		return nil, tokenErr
	}

	offset := pageSize * page

	friendIds, err := u.friendsRepo.GetFriends(userId, offset, pageSize)
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
