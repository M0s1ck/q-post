package transaction

import (
	"context"
	"user-service/internal/repository"

	"gorm.io/gorm"
)

type GormUnitOfWork struct {
	db    *gorm.DB
	txKey repository.TxKeyType
}

func NewGormUnitOfWork(db *gorm.DB, txKey repository.TxKeyType) *GormUnitOfWork {
	return &GormUnitOfWork{db: db, txKey: txKey}
}

func (g GormUnitOfWork) Do(ctx context.Context, fn func(ctx context.Context) error) error {
	return g.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		txCtx := context.WithValue(ctx, g.txKey, tx)
		return fn(txCtx)
	})
}
