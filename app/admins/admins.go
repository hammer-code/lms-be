package admins

import (
	"context"

	"github.com/hammer-code/lms-be/domain"
)

type (
	UsecaseInterface interface {
		GetUsers(ctx context.Context) (users []domain.User, err error)
	}
)
