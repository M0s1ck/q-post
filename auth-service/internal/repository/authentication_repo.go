package repository

import (
	"auth-service/internal/domain/user"
	"context"
	"errors"
	"fmt"

	"github.com/google/uuid"
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

func (repo *AuthenticationRepo) Create(authUser *user.AuthUser) error {
	ctx := context.Background()
	err := gorm.G[user.AuthUser](repo.db).Create(ctx, authUser)

	var pgErr *pgconn.PgError
	if errors.As(err, &pgErr) && pgErr.Code == duplicateErrCode {
		return fmt.Errorf("%w: create username pass: %v", domain.ErrDuplicate, err)
	}

	if err != nil {
		return fmt.Errorf("%w: create username pass: %v", domain.UnhandledDbError, err)
	}

	return nil
}

func (repo *AuthenticationRepo) GetByUsername(username string) (*user.AuthUser, error) {
	ctx := context.Background()
	user, err := gorm.G[user.AuthUser](repo.db).Where("username = ?", username).First(ctx)

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, fmt.Errorf("%w: get by username: %v", domain.ErrNotFound, err)
	}

	if err != nil {
		return nil, fmt.Errorf("%w: get by username: %v", domain.UnhandledDbError, err)
	}

	return &user, nil
}

func (repo *AuthenticationRepo) UpdateRole(userId uuid.UUID, newRole user.UserRole) error {
	ctx := context.Background()

	affected, err := gorm.G[user.AuthUser](repo.db).Where("id = ?", userId).
		Update(ctx, "role", newRole)

	if affected == 0 {
		return fmt.Errorf("%w: in update role", domain.ErrNotFound)
	}

	if err != nil {
		return fmt.Errorf("%w: in update role: %v", domain.UnhandledDbError, err)
	}

	return nil
}
