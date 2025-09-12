package repository

import (
	"context"
	"errors"
	"fmt"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgconn"
	"gorm.io/gorm"
	"user-service/internal/domain"
)

const duplicateErrCode = "23505"

type UserRepo struct {
	db *gorm.DB
}

func NewUserRepo(dbs *gorm.DB) *UserRepo {
	return &UserRepo{db: dbs}
}

func (repo *UserRepo) GetById(id uuid.UUID) (*domain.User, error) {
	ctx := context.Background()
	user, err := gorm.G[domain.User](repo.db).Where("id = ?", id).First(ctx)

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, fmt.Errorf("%w: get by id: %v", domain.ErrNotFound, err)
	}

	if err != nil {
		return nil, fmt.Errorf("%w: get by id: %v", domain.UnhandledDbError, err)
	}

	return &user, nil
}

func (repo *UserRepo) Create(user *domain.User) (uuid.UUID, error) {
	user.Id = uuid.New()
	ctx := context.Background()
	err := gorm.G[domain.User](repo.db).Create(ctx, user)

	var pgErr *pgconn.PgError
	if errors.As(err, &pgErr) && pgErr.Code == duplicateErrCode {
		return uuid.Nil, fmt.Errorf("%w: create user: %v", domain.ErrDuplicate, err)
	}

	if err != nil {
		return uuid.Nil, fmt.Errorf("%w: create user: %v", domain.UnhandledDbError, err)
	}

	return user.Id, nil
}
