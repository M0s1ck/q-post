package relationships

import (
	"fmt"
	"github.com/google/uuid"
	"user-service/internal/usecase"
)

type FollowUseCase struct {
	tokenVal usecase.AccessTokenValidator
}

func (u *FollowUseCase) Follow(followeeId uuid.UUID, token string) error {
	followerId, tokenErr := u.tokenVal.ValidateUserTokenAndGetId(token)
	if tokenErr != nil {
		return tokenErr
	}

	// TODO: consider adding number of friends/followers/followees to user model
	_ = followerId
	panic(fmt.Errorf("implement me"))
}
