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
		Name:         user.Name,
		Description:  user.Description,
		Birthday:     user.Birthday,
		CreatedAt:    user.CreatedAt,
	}

	return &userDto
}

func UserFromCreateRequest(dto *dto.UserToCreate) *domain.User {
	user := domain.User{
		Id:        dto.UserId,
		Username:  dto.Username,
		CreatedAt: time.Now(),
	}

	return &user
}
