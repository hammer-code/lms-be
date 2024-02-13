package db

import (
	"context"

	"github.com/sirupsen/logrus"
)

func (d *dbTX) StartTransaction(ctx context.Context, fn func(txCtx context.Context) error) error {
	tx := d.db.Begin()

	defer tx.Rollback()

	if err := fn(ctx); err != nil {
		logrus.Error("transaction database return err: ", err)
		return err
	}

	if err := tx.Commit().Error; err != nil {
		logrus.Error("error commit database: ", err)
		return err
	}
	return nil
}
