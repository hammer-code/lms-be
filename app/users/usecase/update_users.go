package usecase

import (
	"context"

	"github.com/hammer-code/lms-be/domain"
	"github.com/sirupsen/logrus"
)

func (us *usecase) UpdateProfileUser(ctx context.Context, userReq domain.UserUpdateProfile, id int) error {
	if err := us.dbTX.StartTransaction(ctx, func(txCtx context.Context) error {
		err := us.userRepo.UpdateProfileUser(ctx, userReq, id)
		if err != nil {
			logrus.Error("us.UpdateUser: failed to update users. ", err)
			return err
		}
		return nil
	}); err != nil {
		logrus.Error(err)
		return err
	}
	return nil
}
