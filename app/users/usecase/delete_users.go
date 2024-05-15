package usecase

import (
	"context"

	"github.com/sirupsen/logrus"
)

func (us *usecase) DeleteUser(ctx context.Context, id int8) error {
	_, err := us.GetUserById(ctx, id)
	if err != nil {
		logrus.WithError(err).Error("us.DeleteUser: failed to find user")
		return err
	}

	err = us.userRepo.DeleteUser(ctx, id)
	if err != nil {
		logrus.WithError(err).Error("us.DeleteUser: failed to delete user")
		return err
	}
	return nil
}
