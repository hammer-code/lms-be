package repository

import (
	"context"

	"github.com/hammer-code/lms-be/domain"
	"github.com/sirupsen/logrus"
)

func (repo *repository) UpdateProfileUser(ctx context.Context, userReq domain.UserUpdateProfile, id int) error {
	user:=domain.UserUpdateProfileToUser(userReq)
	err := repo.db.DB(ctx).Where("id = ?", id).Updates(&user).Error
	if err != nil {
		logrus.Error("repo.Update User: failed to update users")
		return err
	}
	return nil
}
