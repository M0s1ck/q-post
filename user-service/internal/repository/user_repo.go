package repository

import (
	"errors"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"user-service/internal/domain"
)

type UserRepo struct {
	db *gorm.DB
}

func NewUserRepo(dbs *gorm.DB) *UserRepo {
	return &UserRepo{db: dbs}
}

func (repo *UserRepo) GetById(id uuid.UUID) (domain.User, error) {
	user := domain.User{}
	res := repo.db.First(&user, id)

	if errors.Is(res.Error, gorm.ErrRecordNotFound) {
		return user, domain.ErrorNotFound
	}

	if res.Error != nil {
		return user, domain.UnhandledDbError
	}

	return user, nil
}
