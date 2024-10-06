package repository

import (
	"context"

	"github.com/hammer-code/lms-be/domain"
	"github.com/sirupsen/logrus"
)

func (repo *repository) UpdateUseImage(ctx context.Context, id uint) error {
	err := repo.db.DB(ctx).Model(&domain.Image{}).
		Where("id = ?", id).Update("is_used", true).Error
	if err != nil {
		logrus.Error("repo.Update: failed to update use image")
		return err
	}
	return nil
}
