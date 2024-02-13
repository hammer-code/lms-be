package db

import (
	"context"

	"github.com/hammer-code/lms-be/domain"
	"gorm.io/gorm"
)

// get db transaction or repo db(without transaction)
func (d *dbTX) DB(ctx context.Context) *gorm.DB {
	tx, ok := ctx.Value(domain.ContextDatabaseTransaction).(*gorm.DB)
	if !ok {
		return d.db
	}
	return tx
}
