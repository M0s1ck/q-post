package mapper

import (
	"user-service/internal/domain"
	"user-service/internal/dto"
)

func GetUserDto(user *domain.User) dto.UserDto {
	userDto := dto.UserDto{
		Id:           user.Id,
		Username:     user.Username,
		PostKarma:    user.PostKarma,
		CommentKarma: user.CommentKarma,
		CreatedAt:    user.CreatedAt,
	}

	return userDto
}
