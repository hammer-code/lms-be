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

func (repo *repository) CreateUser(ctx context.Context, userReq domain.User) (user domain.User, err error) {
	err = repo.db.DB(ctx).Create(&userReq).Error
	if err != nil {
		logrus.Error("repo.CreateUser : failed to create user")
		return user, err
	}
	return userReq, nil
}

func (repo *repository) FindByEmail(ctx context.Context, email string) (user domain.User, err error) {
	err = repo.db.DB(ctx).Where("email = ?", email).First(&user).Error
	if err != nil {
		logrus.Error("repo.FindByEmail: failed to find user")
		return
	}
	return user, nil
}
