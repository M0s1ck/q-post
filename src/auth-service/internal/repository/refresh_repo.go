package repository

import (
	"context"
	"errors"
	"fmt"

	"gorm.io/gorm"

	"auth-service/internal/domain"
	"auth-service/internal/domain/refresh"
)

type RefreshTokenRepo struct {
	db *gorm.DB
}

func NewRefreshTokenRepo(dbs *gorm.DB) *RefreshTokenRepo {
	return &RefreshTokenRepo{db: dbs}
}

func (repo *RefreshTokenRepo) Create(refreshModel *refresh.RefreshToken) error {
	ctx := context.Background()
	err := gorm.G[refresh.RefreshToken](repo.db).Create(ctx, refreshModel)

	if err != nil {
		return fmt.Errorf("%w: create refresh token: %v", domain.UnhandledDbError, err)
	}

	return nil
}

func (repo *RefreshTokenRepo) GetByTokenHash(tokenHash string) (*refresh.RefreshToken, error) {
	ctx := context.Background()
	model, err := gorm.G[refresh.RefreshToken](repo.db).Where("token_hash = ?", tokenHash).First(ctx)

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, fmt.Errorf("%w: get by token hash: %v", domain.ErrNotFound, err)
	}

	if err != nil {
		return nil, fmt.Errorf("%w: get by token hash: %v", domain.UnhandledDbError, err)
	}

	return &model, nil
}

func (repo *RefreshTokenRepo) RemoveByTokenHash(tokenHash string) error {
	ctx := context.Background()
	_, err := gorm.G[refresh.RefreshToken](repo.db).Where("token_hash = ?", tokenHash).Delete(ctx)

	if err != nil {
		return fmt.Errorf("%w: remove by token hash: %v", domain.UnhandledDbError, err)
	}

	return nil
}
