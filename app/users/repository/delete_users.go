package repository

import (
	"context"

	"github.com/hammer-code/lms-be/domain"
	"github.com/sirupsen/logrus"
)

func (repo *repository) DeleteUser(ctx context.Context, id int8) error {
	err := repo.db.DB(ctx).Where("id = ?", id).Delete(&domain.User{}).Error
	if err != nil {
		logrus.Error("repo.Delete: failed to delete user")
		return err
	}
	return nil
}
