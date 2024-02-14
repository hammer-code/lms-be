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

func (us *usecase) Login(ctx context.Context, userReq domain.Login) (user domain.User, token string, err error) {
	err = us.dbTX.StartTransaction(ctx, func(txCtx context.Context) error {
		user, err = us.userRepo.FindByEmail(ctx, userReq.Email)
		if err != nil {
			logrus.Error("us.GetUsers: failed to get users. ", err)
			return err
		}
		return nil
	})

	if err != nil {
		if err != nil {
			logrus.Error("us.GetUsers: failed to get users. ", err)
			return
		}
	}

	signToken, err := us.jwt.GenerateAccessToken(ctx, &user)
	if err != nil {
		logrus.Error("us.Login: failed to login. ", err)
		return
	}

	return user, *signToken, nil
}
