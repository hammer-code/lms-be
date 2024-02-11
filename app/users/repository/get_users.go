package repository

import (
	"context"

	"github.com/hammer-code/lms-be/domain"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

func (repo *repository) GetUsers(ctx context.Context, tx *gorm.DB) (users []domain.User, err error) {
	err = tx.Find(&users).Error
	if err != nil {
		logrus.Error("repo.GetUsers: failed to get users")
		return
	}
	return
}
