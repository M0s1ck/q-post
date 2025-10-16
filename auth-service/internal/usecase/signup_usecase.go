package usecase

import (
	"auth-service/internal/domain/user"
	"github.com/google/uuid"

	"auth-service/internal/dto"
)

type UserCreator interface {
	Create(userId uuid.UUID, username string, pass string, role user.UserRole) error
}

type SignUpUsecase struct {
	userCreator       UserCreator
	refreshTokenSavor RefreshTokenSaver
	accessTokenIssuer AccessTokenIssuer
}

func NewSignUpUsecase(userCreator UserCreator, refreshTokenSavor RefreshTokenSaver,
	accessTokenIssuer AccessTokenIssuer) *SignUpUsecase {
	return &SignUpUsecase{
		userCreator:       userCreator,
		refreshTokenSavor: refreshTokenSavor,
		accessTokenIssuer: accessTokenIssuer,
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

	refreshErr := uc.refreshTokenSavor.GenerateNewAndSave(userId)

	if refreshErr != nil {
		return nil, refreshErr
	}

	accessToken, tokenErr := uc.accessTokenIssuer.IssueAccessToken(userId, username, userRole)

	if tokenErr != nil {
		return nil, tokenErr
	}

	return &dto.UserIdAndTokens{UserId: userId, AccessToken: accessToken}, nil
}
