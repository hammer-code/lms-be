package users

import (
	"context"

	"github.com/hammer-code/lms-be/domain"
)

type (
	UserRepository interface {
		GetUsers(ctx context.Context) (users []domain.User, err error)
	}

	UserUsecase interface {
		GetUsers(ctx context.Context) (users []domain.User, err error)
	}
)
