package usecase

import (
	"github.com/google/uuid"

	"auth-service/internal/domain"
	"auth-service/internal/dto"
)

type UserCreater interface {
	Create(user *domain.AuthUser) error
}

type SignUpUsecase struct {
	repo        UserCreater
	tokenIssuer TokenIssuer
	passHasher  PasswordHasher
}

func NewSignUpUsecase(rep UserCreater, tokenIssuer TokenIssuer, hasher PasswordHasher) *SignUpUsecase {
	return &SignUpUsecase{
		repo:        rep,
		tokenIssuer: tokenIssuer,
		passHasher:  hasher,
	}
}

func (uc *SignUpUsecase) SignUpWithUsername(usPass *dto.UsernamePass) (*dto.UserIdAndTokens, error) {
	passHash, err := uc.passHasher.Hash(usPass.Password)
	if err != nil {
		return nil, err
	}

	id := uuid.New() // TODO: think through

	authUser := domain.AuthUser{
		Id:             id,
		Username:       usPass.Username,
		HashedPassword: passHash,
		Role:           domain.RoleUser,
	}

	dbErr := uc.repo.Create(&authUser)

	if dbErr != nil {
		return nil, dbErr
	}

	// TODO: Call here to user-service

	accessToken, tokenErr := uc.tokenIssuer.IssueAccessToken(authUser.Id, authUser.Username, authUser.Role)

	if tokenErr != nil {
		return nil, tokenErr
	}

	return &dto.UserIdAndTokens{UserId: id, AccessToken: accessToken}, nil
}
