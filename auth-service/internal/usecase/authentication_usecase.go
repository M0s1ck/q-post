package usecase

import (
	"auth-service/internal/domain"
	"github.com/google/uuid"

	"auth-service/internal/dto"
	"auth-service/internal/repository"
)

type PasswordHasher interface {
	Hash(password string) (string, error)
	Verify(password string, hash string) (bool, error)
}

type TokenIssuer interface {
	CreateAccessToken(username string) (string, error)
}

type AuthenticationUsecase struct {
	repo        *repository.AuthenticationRepo
	passHasher  PasswordHasher
	tokenIssuer TokenIssuer
}

func NewAuthenticationUsecase(rep *repository.AuthenticationRepo, hasher PasswordHasher, tokenIssuer TokenIssuer) *AuthenticationUsecase {
	return &AuthenticationUsecase{
		repo:        rep,
		passHasher:  hasher,
		tokenIssuer: tokenIssuer,
	}
}

func (uc *AuthenticationUsecase) SignUp(usPass *dto.UsernamePass) (*dto.UserIdAndTokens, error) {
	passHash, err := uc.passHasher.Hash(usPass.Password)
	if err != nil {
		return nil, err
	}

	id := uuid.New() // TODO: think through

	authUser := domain.AuthUser{
		Id:             id,
		Username:       usPass.Username,
		HashedPassword: passHash,
	}

	dbErr := uc.repo.Create(&authUser)

	if dbErr != nil {
		return nil, dbErr
	}

	accessToken, tokenErr := uc.tokenIssuer.CreateAccessToken(usPass.Username)

	if tokenErr != nil {
		return nil, tokenErr
	}

	return &dto.UserIdAndTokens{UserId: id, AccessToken: accessToken}, nil
}
