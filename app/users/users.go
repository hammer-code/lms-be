package users

import (
	"context"
	"time"

	"github.com/hammer-code/lms-be/domain"
)

type (
	UserRepository interface {
		GetUsers(ctx context.Context) (users []domain.User, err error)
		CreateUser(ctx context.Context, userReq domain.User) (user domain.User, err error)
		FindByEmail(ctx context.Context, email string) (user domain.User, err error)
		LogoutUser(ctx context.Context, token string, expiredAt time.Time) error
		ExpiredToken(ctx context.Context, token string) error
	}
	UserUsecase interface {
		GetUsers(ctx context.Context) (users []domain.User, err error)
		Register(ctx context.Context, userReq domain.User) (user domain.User, err error)
		Login(ctx context.Context, userReq domain.Login) (user domain.User, token string, err error)
		Logout(ctx context.Context, token string) error
	}
)
