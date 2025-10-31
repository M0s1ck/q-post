package mapper

import (
	"fmt"
	"time"

	"user-service/internal/domain/user"
	"user-service/internal/dto"
)

const dateLayout = "2006-01-02"

func GetUserDto(user *user.User) *dto.UserResponse {
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

func UserFromCreateRequest(dto *dto.UserToCreate) *user.User {
	us := user.User{
		Id:        dto.UserId,
		Username:  dto.Username,
		CreatedAt: time.Now(),
	}

	return &us
}

func GetUserDetailsFromDto(usDetDto *dto.UserDetailStr) (*user.UserDetails, error) {
	details := user.UserDetails{
		Name:        usDetDto.Name,
		Description: usDetDto.Description,
	}

	if usDetDto.Birthday != nil {
		bday, dateErr := time.Parse(dateLayout, *usDetDto.Birthday)
		if dateErr != nil {
			return nil, fmt.Errorf("user details dto: expected date format 2006-01-02: %v", dateErr)
		}
		details.Birthday = &bday
	}

	return &details, nil
}

func GetUserSummaries(users []user.User) []dto.UserSummary {
	var sums = make([]dto.UserSummary, len(users))

	for i, us := range users {
		sum := GetUserSummary(&us)
		sums[i] = *sum
	}

	return sums
}

func GetUserSummary(us *user.User) *dto.UserSummary {
	return &dto.UserSummary{
		Id:       us.Id,
		Username: us.Username,
	}
}
