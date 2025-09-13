package usecase

import (
	"github.com/google/uuid"
	"user-service/internal/dto"
	"user-service/internal/mapper"
	"user-service/internal/repository"
)

type UserUseCase struct {
	userRepo *repository.UserRepo
}

func NewUserUseCase(userRepo *repository.UserRepo) *UserUseCase {
	return &UserUseCase{userRepo: userRepo}
}

func (useCase *UserUseCase) GetById(id uuid.UUID) (*dto.UserResponse, error) {
	user, err := useCase.userRepo.GetById(id)

	if err != nil {
		return nil, err
	}

	userDto := mapper.GetUserDto(user)
	return userDto, err
}

func (useCase *UserUseCase) Create(userDto *dto.UserToCreate) (*dto.UuidOnlyResponse, error) {
	user := mapper.UserFromCreateRequest(userDto)
	id, err := useCase.userRepo.Create(user)
	return &dto.UuidOnlyResponse{Id: id}, err
}

func (useCase *UserUseCase) UpdateDetails(id uuid.UUID, details *dto.UserDetailsToUpdate) error {
	err := useCase.userRepo.UpdateDetails(id, details)
	return err
}

func (useCase *UserUseCase) Delete(id uuid.UUID) error {
	err := useCase.userRepo.Delete(id)
	return err
}
