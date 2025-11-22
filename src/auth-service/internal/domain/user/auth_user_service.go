package user

import (
	"github.com/google/uuid"

	"auth-service/internal/domain"
)

type AuthUserService struct {
	passHasher HasherVerifier
	repo       AuthUserCreatorGetter
}

func NewAuthUserService(repo AuthUserCreatorGetter, passHasher HasherVerifier) *AuthUserService {
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

func (serv *AuthUserService) GetVerifiedByUsername(username string, pass string) (*AuthUser, error) {
	us, err := serv.repo.GetByUsername(username)

	if err != nil {
		return nil, err
	}

	valid, hashErr := serv.passHasher.Verify(pass, us.HashedPassword)

	if hashErr != nil {
		return nil, hashErr
	}

	if !valid {
		return nil, domain.ErrWrongPassword
	}

	return us, nil
}
