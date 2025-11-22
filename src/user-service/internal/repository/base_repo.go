package repository

import (
	"context"

	"gorm.io/gorm"
)

type TxKeyType struct{}

// TODO: maybe move to separate repo/gorm pkg

type BaseRepo struct {
	db    *gorm.DB
	txKey TxKeyType
}

func NewBaseRepo(db *gorm.DB, txKey TxKeyType) *BaseRepo {
	return &BaseRepo{db: db, txKey: txKey}
}

// getTx get database client: global or transactional if we're in transaction
func (rep *BaseRepo) getTx(ctx context.Context) *gorm.DB {
	tx, ok := ctx.Value(rep.txKey).(*gorm.DB)
	if ok {
		return tx
	}
	return rep.db
}
