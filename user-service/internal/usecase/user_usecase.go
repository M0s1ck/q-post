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

func (useCase *UserUseCase) GetById(id uuid.UUID) (dto.UserDto, error) {
	user, err := useCase.userRepo.GetById(id)
	userDto := mapper.GetUserDto(&user)
	return userDto, err
}
