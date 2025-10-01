package repository

import (
	"auth-service/internal/domain"
	"context"
	"fmt"
	"gorm.io/gorm"
)

type RefreshTokenRepo struct {
	db *gorm.DB
}

func NewRefreshTokenRepo(dbs *gorm.DB) *RefreshTokenRepo {
	return &RefreshTokenRepo{db: dbs}
}

func (repo *RefreshTokenRepo) Create(refreshModel *domain.RefreshToken) error {
	ctx := context.Background()
	err := gorm.G[domain.RefreshToken](repo.db).Create(ctx, refreshModel)

	if err != nil {
		return fmt.Errorf("%w: create refresh token: %v", domain.UnhandledDbError, err)
	}

	return nil
}
