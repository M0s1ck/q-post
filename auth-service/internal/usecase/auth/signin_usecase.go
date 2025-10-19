package auth

import (
	"auth-service/internal/domain/user"
	"auth-service/internal/dto"
	"auth-service/internal/usecase"
)

type UserVerifier interface {
	// GetVerifiedByUsername Gets user by username verified by password;
	// domain.ErrWrongPassword if pass is wrong
	GetVerifiedByUsername(username string, pass string) (*user.AuthUser, error)
}

type SignInUsecase struct {
	userVerifier          UserVerifier
	refreshTokenGenerator RefreshTokenGenerator
	tokenIssuer           usecase.AccessTokenIssuer
}

func NewSignInUsecase(verifier UserVerifier, refreshTokenGenerator RefreshTokenGenerator, tokenIssuer usecase.AccessTokenIssuer) *SignInUsecase {
	return &SignInUsecase{
		userVerifier:          verifier,
		refreshTokenGenerator: refreshTokenGenerator,
		tokenIssuer:           tokenIssuer,
	}
}

func (uc *SignInUsecase) SignInByUsername(usPass *dto.UsernamePass) (*dto.UserIdAndTokens, error) {
	us, err := uc.userVerifier.GetVerifiedByUsername(usPass.Username, usPass.Password)

	if err != nil {
		return nil, err
	}

	accessToken, tokenErr := uc.tokenIssuer.IssueAccessToken(us.Id, us.Username, us.Role)

	if tokenErr != nil {
		return nil, tokenErr
	}

	refresh, refreshErr := uc.refreshTokenGenerator.GenerateNewAndSave(us.Id)

	if refreshErr != nil {
		return nil, refreshErr
	}

	response := dto.UserIdAndTokens{
		UserId:       us.Id,
		AccessToken:  accessToken,
		RefreshToken: refresh,
	}

	return &response, nil
}
