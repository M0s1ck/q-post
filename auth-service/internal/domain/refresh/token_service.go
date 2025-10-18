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

func (serv *TokenService) GenerateNewAndSave(userId uuid.UUID) (uuid.UUID, error) {
	token := uuid.New()
	tokenHash, hashErr := serv.hasher.Hash(token.String())

	if hashErr != nil {
		return uuid.Nil, hashErr
	}

	tokenModel := NewRefreshToken(tokenHash, userId)
	dbErr := serv.repo.Create(&tokenModel)

	if dbErr != nil {
		return uuid.Nil, dbErr
	}

	return token, nil
}
