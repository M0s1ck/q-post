package usecase

import (
	"github.com/google/uuid"

	"auth-service/internal/dto"
	"auth-service/internal/repository"
)

type PasswordHasher interface {
	Hash(password string) (string, error)
	Verify(password string, hash string) (bool, error)
}

type AuthenticationUsecase struct {
	repo       *repository.AuthenticationRepo
	passHasher PasswordHasher
}

func NewAuthenticationUsecase(rep *repository.AuthenticationRepo, hasher PasswordHasher) *AuthenticationUsecase {
	return &AuthenticationUsecase{
		repo:       rep,
		passHasher: hasher,
	}
}

func (uc *AuthenticationUsecase) SignUp(usPass *dto.UsernamePass) (uuid.UUID, error) {
	username := usPass.Username
	passHash, err := uc.passHasher.Hash(usPass.Password)

	if err != nil {
		return uuid.Nil, err
	}

	id := uuid.New() // TODO: think through

	dbErr := uc.repo.CreateByUsernamePass(id, username, passHash)
	return id, dbErr
}
