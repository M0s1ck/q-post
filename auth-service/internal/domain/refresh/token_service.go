package refresh

import (
	"github.com/google/uuid"
)

type TokenService struct {
	repo   TokenSaver
	hasher Hasher
}

func NewRefreshTokenService(repo TokenSaver, hasher Hasher) *TokenService {
	return &TokenService{
		repo:   repo,
		hasher: hasher,
	}
}

func (serv *TokenService) GenerateNewAndSave(userId uuid.UUID) error {
	token := uuid.New()
	tokenHash, hashErr := serv.hasher.Hash(token.String())

	if hashErr != nil {
		return hashErr
	}

	tokenModel := NewRefreshToken(tokenHash, userId)

	dbErr := serv.repo.Create(&tokenModel)
	return dbErr
}
