package repository

import (
	"context"
	"errors"
	"fmt"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgconn"
	"gorm.io/gorm"

	"user-service/internal/domain"
	"user-service/internal/domain/user"
)

const duplicateErrCode = "23505"

type UserRepo struct {
	db *gorm.DB
}

func NewUserRepo(dbs *gorm.DB) *UserRepo {
	return &UserRepo{db: dbs}
}

func (repo *UserRepo) GetById(id uuid.UUID) (*user.User, error) {
	ctx := context.Background()
	user, err := gorm.G[user.User](repo.db).Where("id = ?", id).First(ctx)

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, fmt.Errorf("%w: get by id: %v", domain.ErrNotFound, err)
	}

	if err != nil {
		return nil, fmt.Errorf("%w: get by id: %v", domain.UnhandledDbError, err)
	}

	return &user, nil
}

func (repo *UserRepo) Create(us *user.User) error {
	ctx := context.Background()
	err := gorm.G[user.User](repo.db).Create(ctx, us)

	var pgErr *pgconn.PgError
	if errors.As(err, &pgErr) && pgErr.Code == duplicateErrCode {
		return fmt.Errorf("%w: create user: %v", domain.ErrDuplicate, err)
	}

	if err != nil {
		return fmt.Errorf("%w: create user: %v", domain.UnhandledDbError, err)
	}

	return nil
}

func (repo *UserRepo) UpdateDetails(id uuid.UUID, details *user.UserDetails) error {
	ctx := context.Background()

	affected, err := gorm.G[user.User](repo.db).Where("id = ?", id).
		Updates(ctx, user.User{Name: details.Name, Description: details.Description, Birthday: details.Birthday})

	if affected == 0 {
		return fmt.Errorf("%w: update details", domain.ErrNotFound)
	}

	if err != nil {
		return fmt.Errorf("%w: update details: %v", domain.UnhandledDbError, err)
	}

	return nil
}

func (repo *UserRepo) Delete(id uuid.UUID) error {
	ctx := context.Background()
	affected, err := gorm.G[user.User](repo.db).Where("id = ?", id).Delete(ctx)

	if affected == 0 {
		return fmt.Errorf("%w: delete", domain.ErrNotFound)
	}

	if err != nil {
		return fmt.Errorf("%w: delete: %v", domain.UnhandledDbError, err)
	}

	return nil
}
