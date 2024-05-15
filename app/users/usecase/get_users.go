package usecase

import (
	"context"

	"github.com/hammer-code/lms-be/domain"
	"github.com/sirupsen/logrus"
)

func (us *usecase) GetUsers(ctx context.Context) (users []domain.User, err error) {
	if err = us.dbTX.StartTransaction(ctx, func(txCtx context.Context) error {
		users, err = us.userRepo.GetUsers(txCtx)
		if err != nil {
			logrus.Error("us.GetUsers: failed to get users. ", err)
			return err
		}
		return nil
	}); err != nil {
		logrus.Error(err)
		return
	}
	return
}

func (us *usecase) GetUserById(ctx context.Context, userID int8) (user domain.User, err error) {
	if err = us.dbTX.StartTransaction(ctx, func(txCtx context.Context) error {
		user, err = us.userRepo.FindById(txCtx, userID)
		if err != nil {
			logrus.Error("us.GetUserById: failed to get user. ", err)
			return err
		}
		return nil
	}); err != nil {
		logrus.Error(err)
		return
	}
	return
}
