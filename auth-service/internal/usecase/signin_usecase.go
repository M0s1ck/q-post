package usecase

import (
	"auth-service/internal/domain"
	"auth-service/internal/dto"
)

type UserGetter interface {
	GetByUsername(username string) (*domain.AuthUser, error)
}

type SignInUsecase struct {
	repo        UserGetter
	tokenIssuer TokenIssuer
	passHasher  PasswordHasher
}

func NewSignInUsecase(repo UserGetter, tokenIssuer TokenIssuer, passHasher PasswordHasher) *SignInUsecase {
	return &SignInUsecase{
		repo:        repo,
		tokenIssuer: tokenIssuer,
		passHasher:  passHasher,
	}
}

func (uc *SignInUsecase) SignInByUsername(usPass *dto.UsernamePass) (*dto.UserIdAndTokens, error) {
	user, err := uc.repo.GetByUsername(usPass.Username)

	if err != nil {
		return nil, err
	}

	valid, hashErr := uc.passHasher.Verify(usPass.Password, user.HashedPassword)

	if hashErr != nil {
		return nil, hashErr
	}

	if !valid {
		return nil, domain.ErrWrongPassword
	}

	accessToken, tokenErr := uc.tokenIssuer.CreateAccessToken(user.Id, user.Username, user.Role)

	if tokenErr != nil {
		return nil, tokenErr
	}

	response := dto.UserIdAndTokens{
		UserId:      user.Id,
		AccessToken: accessToken,
	}

	return &response, nil
}
