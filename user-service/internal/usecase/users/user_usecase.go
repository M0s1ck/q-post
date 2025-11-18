package users

import (
	"github.com/google/uuid"

	"user-service/internal/domain/user"
	"user-service/internal/dto"
	"user-service/internal/mapper"
	"user-service/internal/usecase"
)

const authServiceIssuer = "auth-service"

type UserRepo interface {
	GetById(id uuid.UUID) (*user.User, error)
	Create(us *user.User) error
	UpdateDetails(id uuid.UUID, details *user.UserDetails) error
	Delete(id uuid.UUID) error
}

// UserUseCase
// Basic operations with users
type UserUseCase struct {
	userRepo             UserRepo
	accessTokenValidator usecase.AccessTokenValidator
}

func NewUserUseCase(userRepo UserRepo, jwtValidator usecase.AccessTokenValidator) *UserUseCase {
	return &UserUseCase{
		userRepo:             userRepo,
		accessTokenValidator: jwtValidator,
	}
}

func (u *UserUseCase) GetById(id uuid.UUID) (*dto.UserResponse, error) {
	us, err := u.userRepo.GetById(id)
	if err != nil {
		return nil, err
	}

	userDto := mapper.GetUserDto(us)
	return userDto, err
}

func (u *UserUseCase) GetMe(token string) (*dto.UserResponse, error) {
	userId, tokenErr := u.accessTokenValidator.ValidateUserTokenAndGetId(token)
	if tokenErr != nil {
		return nil, tokenErr
	}

	return u.GetById(userId)
}

func (u *UserUseCase) Create(userDto *dto.UserToCreate, token string) (*dto.UuidOnlyResponse, error) {
	tokenErr := u.accessTokenValidator.ValidateApiTokenIssuedAt(token, authServiceIssuer)
	if tokenErr != nil {
		return nil, tokenErr
	}

	us := mapper.UserFromCreateRequest(userDto)
	err := u.userRepo.Create(us)
	return &dto.UuidOnlyResponse{Id: us.Id}, err
}

func (u *UserUseCase) UpdateDetails(userDetailsDto *dto.UserDetailStr, token string) error {
	details, dtoErr := mapper.GetUserDetailsFromDto(userDetailsDto)
	if dtoErr != nil {
		return dtoErr
	}

	id, tokenErr := u.accessTokenValidator.ValidateUserTokenAndGetId(token)
	if tokenErr != nil {
		return tokenErr
	}

	err := u.userRepo.UpdateDetails(id, details)
	return err
}

func (u *UserUseCase) Delete(id uuid.UUID, token string) error {
	tokenErr := u.accessTokenValidator.ValidateUserTokenBySubIdOrRole(token, id, user.RoleAdmin)
	if tokenErr != nil {
		return tokenErr
	}

	err := u.userRepo.Delete(id)
	return err
}
