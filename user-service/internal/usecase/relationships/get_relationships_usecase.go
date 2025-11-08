package relationships

import (
	"errors"

	"github.com/google/uuid"

	"user-service/internal/domain"
	"user-service/internal/dto"
	"user-service/internal/mapper"
	"user-service/internal/usecase"
)

const DefaultPageSize = 20

const (
	friendshipStatus = "friend"
	followerStatus   = "follower"
	followeeStatus   = "followee"
	nobodyStatus     = "nobody"
)

type GetRelationshipsUseCase struct {
	relationshipRepo relationshipsGetter
	userRepo         userGetter
	tokenVal         usecase.AccessTokenValidator
}

func NewGetRelationshipsUseCase(fRepo relationshipsGetter, uRepo userGetter, tVal usecase.AccessTokenValidator) *GetRelationshipsUseCase {
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

func (u *GetRelationshipsUseCase) GetRelationshipStatus(userId uuid.UUID, token string) (*dto.RelationshipStatus, error) {
	senderId, tokenErr := u.tokenVal.ValidateUserTokenAndGetId(token)
	if tokenErr != nil {
		return nil, tokenErr
	}

	rel, err := u.relationshipRepo.GetRelationship(senderId, userId)
	if errors.Is(err, domain.ErrNotFound) {
		return &dto.RelationshipStatus{Status: nobodyStatus}, nil
	}

	if rel.AreFriends {
		return &dto.RelationshipStatus{Status: friendshipStatus}, nil
	}

	if userId == rel.FolloweeId {
		return &dto.RelationshipStatus{Status: followeeStatus}, nil
	}

	return &dto.RelationshipStatus{Status: followerStatus}, nil
}
