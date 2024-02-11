package users

import (
	"context"

	"github.com/hammer-code/lms-be/domain"
	"gorm.io/gorm"
)

type (
	UserRepository interface {
		GetUsers(ctx context.Context, tx *gorm.DB) (users []domain.User, err error)
	}
)
