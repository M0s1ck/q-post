package repository

import (
	"context"

	"gorm.io/gorm"
)

// TODO: maybe move to separate repo/gorm pkg

type BaseRepo struct {
	db *gorm.DB
}

func NewBaseRepo(db *gorm.DB) *BaseRepo {
	return &BaseRepo{db: db}
}

// getTx get database client: global or transactional if we're in transaction
func (rep *BaseRepo) getTx(ctx context.Context) *gorm.DB {
	tx, ok := ctx.Value("tx").(*gorm.DB)
	if ok {
		return tx
	}
	return rep.db
}
