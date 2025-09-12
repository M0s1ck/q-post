package mapper

import (
	"time"
	"user-service/internal/domain"
	"user-service/internal/dto"
)

func GetUserDto(user *domain.User) *dto.UserResponse {
	userDto := dto.UserResponse{
		Id:           user.Id,
		Username:     user.Username,
		PostKarma:    user.PostKarma,
		CommentKarma: user.CommentKarma,
		CreatedAt:    user.CreatedAt,
	}

	return &userDto
}

func UserFromCreateRequest(dto *dto.UserToCreate) *domain.User {
	user := domain.User{
		Username:  dto.Username,
		CreatedAt: time.Now(),
	}

	return &user
}
