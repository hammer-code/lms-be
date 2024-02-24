package repository

import (
	"context"

	"github.com/hammer-code/lms-be/domain"
	"github.com/sirupsen/logrus"
)

// query untuk get user di database
func (repo *repository) GetUsers(ctx context.Context) (users []domain.User, err error) {
	err = repo.db.DB(ctx).Find(&users).Error
	if err != nil {
		logrus.Error("repo.GetUsers: failed to get users")
		return
	}
	return
}
