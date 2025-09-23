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
	CreateAccessToken(id uuid.UUID, username string, role domain.UserRole) (string, error)
}

type SignUpUsecase struct {
	repo        *repository.AuthenticationRepo
	passHasher  PasswordHasher
	tokenIssuer TokenIssuer
}

func NewSignUpUsecase(rep *repository.AuthenticationRepo, hasher PasswordHasher, tokenIssuer TokenIssuer) *SignUpUsecase {
	return &SignUpUsecase{
		repo:        rep,
		passHasher:  hasher,
		tokenIssuer: tokenIssuer,
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

	accessToken, tokenErr := uc.tokenIssuer.CreateAccessToken(authUser.Id, authUser.Username, authUser.Role)

	if tokenErr != nil {
		return nil, tokenErr
	}

	return &dto.UserIdAndTokens{UserId: id, AccessToken: accessToken}, nil
}
