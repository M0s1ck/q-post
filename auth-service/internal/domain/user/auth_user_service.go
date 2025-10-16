package user

import "github.com/google/uuid"

type AuthUserService struct {
	passHasher PassHasher
	repo       AuthUserCreator
}

func NewAuthUserService(repo AuthUserCreator, passHasher PassHasher) *AuthUserService {
	return &AuthUserService{passHasher: passHasher, repo: repo}
}

func (serv *AuthUserService) Create(userId uuid.UUID, username string, pass string, role UserRole) error {
	passHash, hashErr := serv.passHasher.Hash(pass)
	if hashErr != nil {
		return nil
	}

	authUser := AuthUser{
		Id:             userId,
		Username:       username,
		HashedPassword: passHash,
		Role:           role,
	}

	dbErr := serv.repo.Create(&authUser)
	return dbErr
}
