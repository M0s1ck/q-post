package auth

import (
	"github.com/google/uuid"

	"auth-service/internal/domain/user"
	"auth-service/internal/dto"
	"auth-service/internal/usecase"
)

type UserCreator interface {
	Create(userId uuid.UUID, username string, pass string, role user.UserRole) error
}

type SignUpUsecase struct {
	userCreator           UserCreator
	refreshTokenGenerator RefreshTokenGenerator
	accessTokenIssuer     usecase.AccessTokenIssuer
}

func NewSignUpUsecase(userCreator UserCreator, refreshTokenGenerator RefreshTokenGenerator,
	accessTokenIssuer usecase.AccessTokenIssuer) *SignUpUsecase {
	return &SignUpUsecase{
		userCreator:           userCreator,
		refreshTokenGenerator: refreshTokenGenerator,
		accessTokenIssuer:     accessTokenIssuer,
	}
}

func (uc *SignUpUsecase) SignUpWithUsername(usPass *dto.UsernamePass) (*dto.UserIdAndTokens, error) {
	userId := uuid.New()
	username := usPass.Username
	pass := usPass.Password
	userRole := user.RoleUser

	createErr := uc.userCreator.Create(userId, username, pass, userRole)

	if createErr != nil {
		return nil, createErr
	}

	// TODO: Call here to user-service

	refresh, refreshErr := uc.refreshTokenGenerator.GenerateNewAndSave(userId)

	if refreshErr != nil {
		return nil, refreshErr
	}

	accessToken, tokenErr := uc.accessTokenIssuer.IssueAccessToken(userId, username, userRole)

	if tokenErr != nil {
		return nil, tokenErr
	}

	response := dto.UserIdAndTokens{
		UserId:       userId,
		AccessToken:  accessToken,
		RefreshToken: refresh,
	}

	return &response, nil
}
