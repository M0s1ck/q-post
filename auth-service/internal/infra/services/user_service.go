package services

import "auth-service/internal/dto"

type UserServiceClient struct {
	domain string
}

func (client *UserServiceClient) CreateUser(us *dto.UserToCreate) {

}
