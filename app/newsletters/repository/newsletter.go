package repository

import (
	"context"
	"errors"

	"github.com/hammer-code/lms-be/domain"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

func (repo *repository) Subscribe(ctx context.Context, n domain.Newsletter) error {
	if err := repo.db.DB(ctx).Save(&n).Error; err != nil {
		logrus.Error("repo.Update User: failed to update users")
		return err
	}

	return nil
}

func (repo *repository) GetByEmail(ctx context.Context, email string) (*domain.Newsletter, error) {
	n := domain.Newsletter{}

	if err := repo.db.DB(ctx).First(&n, "email = ?", email).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		logrus.Error("repo.Update User: failed to update users")
		return nil, err
	}

	return &n, nil
}
