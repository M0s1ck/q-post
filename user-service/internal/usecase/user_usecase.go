package usecase

import (
	"fmt"

	"github.com/google/uuid"

	"user-service/internal/domain"
	"user-service/internal/dto"
	"user-service/internal/mapper"
	"user-service/internal/repository"
)

const authServiceIssuer = "auth-service"

type UserUseCase struct {
	userRepo             *repository.UserRepo
	accessTokenValidator AccessTokenApiValidator
}

func NewUserUseCase(userRepo *repository.UserRepo, jwtValidator AccessTokenApiValidator) *UserUseCase {
	return &UserUseCase{
		userRepo:             userRepo,
		accessTokenValidator: jwtValidator,
	}
}

func (u *UserUseCase) GetById(id uuid.UUID) (*dto.UserResponse, error) {
	user, err := u.userRepo.GetById(id)

	if err != nil {
		return nil, err
	}

	userDto := mapper.GetUserDto(user)
	return userDto, err
}

func (u *UserUseCase) Create(userDto *dto.UserToCreate, token string) (*dto.UuidOnlyResponse, error) {
	tokenErr := u.accessTokenValidator.ValidateTokenIssuedAt(token, authServiceIssuer) // TODO: test
	if tokenErr != nil {
		return nil, fmt.Errorf("%w: %v", domain.ErrInvalidToken, tokenErr)
	}

	user := mapper.UserFromCreateRequest(userDto)
	err := u.userRepo.Create(user)
	return &dto.UuidOnlyResponse{Id: user.Id}, err
}

func (u *UserUseCase) UpdateDetails(id uuid.UUID, details *dto.UserDetailsToUpdate) error {
	err := u.userRepo.UpdateDetails(id, details)
	return err
}

func (u *UserUseCase) Delete(id uuid.UUID) error {
	err := u.userRepo.Delete(id)
	return err
}
