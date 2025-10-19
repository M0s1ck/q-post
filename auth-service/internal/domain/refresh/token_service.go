package refresh

import (
	"fmt"
	"time"

	"github.com/google/uuid"

	"auth-service/internal/domain"
)

type TokenService struct {
	repo   TokenRepo
	hasher Hasher
}

func NewRefreshTokenService(repo TokenRepo, hasher Hasher) *TokenService {
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

func (serv *TokenService) Verify(token uuid.UUID) (userId uuid.UUID, err error) {
	tokenHash, hashErr := serv.hasher.Hash(token.String())

	if hashErr != nil {
		return uuid.Nil, hashErr
	}

	refreshModel, dbErr := serv.repo.GetByTokenHash(tokenHash)

	if dbErr != nil {
		return uuid.Nil, dbErr
	}

	// Check if it's not expired
	if time.Now().Before(refreshModel.ExpiresAt) {
		return refreshModel.UserId, nil
	}

	// Remove if it's expired
	_ = serv.repo.RemoveByTokenHash(tokenHash)
	return uuid.Nil, fmt.Errorf("refresh token is expired: %v", domain.ErrInvalidToken)
}
