package repository

import (
	"context"

	"github.com/hammer-code/lms-be/domain"
	"github.com/sirupsen/logrus"
)

func (repo *repository) Store(ctx context.Context, data domain.Image) (uint, error) {
	err := repo.db.DB(ctx).Create(&data).Error
	if err != nil {
		logrus.Error("repo.Store : failed to store")
		return data.ID, err
	}
	return data.ID, nil
}
