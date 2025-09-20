package repository

import (
	"context"
	"errors"
	"fmt"

	"github.com/jackc/pgx/v5/pgconn"
	"gorm.io/gorm"

	"auth-service/internal/domain"
)

const duplicateErrCode = "23505"

type AuthenticationRepo struct {
	db *gorm.DB
}

func NewAuthenticationRepo(dbs *gorm.DB) *AuthenticationRepo {
	return &AuthenticationRepo{db: dbs}
}

func (repo *AuthenticationRepo) Create(authUser *domain.AuthUser) error {
	ctx := context.Background()
	err := gorm.G[domain.AuthUser](repo.db).Create(ctx, authUser)

	var pgErr *pgconn.PgError
	if errors.As(err, &pgErr) && pgErr.Code == duplicateErrCode {
		return fmt.Errorf("%w: create username pass: %v", domain.ErrDuplicate, err)
	}

	if err != nil {
		return fmt.Errorf("%w: create username pass: %v", domain.UnhandledDbError, err)
	}

	return nil
}
