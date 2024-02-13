package db

import (
	"context"

	"gorm.io/gorm"
)

type (
	dbTX struct {
		db *gorm.DB
	}

	DatabaseTransaction interface {
		StartTransaction(ctx context.Context, fn func(txCtx context.Context) error) error
	}
)

func NewDBTransaction(db *gorm.DB) DatabaseTransaction {
	return &dbTX{
		db: db,
	}
}
