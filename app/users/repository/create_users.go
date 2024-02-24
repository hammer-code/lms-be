package repository

import (
	"context"
	"github.com/hammer-code/lms-be/domain"
	"github.com/sirupsen/logrus"
)

func (repo *repository) CreateUser(ctx context.Context, userReq domain.User) (user domain.User, err error) {
	err = repo.db.DB(ctx).Create(&userReq).Error
	if err != nil {
		logrus.Error("repo.CreateUser : failed to create user")
		return user, err
	}
	return userReq, nil
}
