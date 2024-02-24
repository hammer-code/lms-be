package usecase

import (
	"context"
	"github.com/hammer-code/lms-be/domain"
	"github.com/sirupsen/logrus"
)

func (us *usecase) Register(ctx context.Context, userReq domain.User) (user domain.User, err error) {
	if err = us.dbTX.StartTransaction(ctx, func(txCtx context.Context) error {
		user, err = us.userRepo.CreateUser(ctx, userReq)
		if err != nil {
			logrus.Error("us.Register: failed to get users. ", err)
			return err
		}
		return nil

	}); err != nil {
		logrus.Error("us.Register: failed to get users. ", err)
		return domain.User{}, err
	}
	return user, nil
}
