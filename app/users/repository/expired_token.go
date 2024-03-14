package repository

import (
	"context"
	"github.com/hammer-code/lms-be/domain"
	"time"
)

func (repo *repository) ExpiredToken(ctx context.Context, token string) error {
	tokenData := &domain.LogoutToken{}
	if err := repo.db.DB(ctx).Model(&domain.LogoutToken{}).Where("token = ?", token).First(&tokenData).Error; err != nil {
		return err
	}

	if tokenData.ExpiredAt.Before(time.Now()) {
		if err := repo.db.DB(ctx).Delete(&tokenData).Error; err != nil {
			return err
		}
	}

	return nil
}
