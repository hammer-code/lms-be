package repository

import (
	"context"
	"github.com/hammer-code/lms-be/domain"
	"time"
)

func (repo *repository) LogoutUser(ctx context.Context, token string, expiredAt time.Time) error {
	err := repo.db.DB(ctx).Create(domain.LogoutToken{
		Token:     token,
		ExpiredAt: expiredAt,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}).Error

	if err != nil {
		return err
	}

	return nil
}
