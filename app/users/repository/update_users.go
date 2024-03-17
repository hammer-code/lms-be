package repository

import (
	"context"

	"github.com/hammer-code/lms-be/domain"
	"github.com/sirupsen/logrus"
)

func (repo *repository) UpdateProfileUser(ctx context.Context, userReq domain.UserUpdateProfile, id int) error {
	err := repo.db.DB(ctx).Where("id = ?", id).Updates(&userReq).Error
	if err != nil {
		logrus.Error("repo.Update User: failed to update users")
		return err
	}
	return nil
}
