package repository

import (
	"context"

	"github.com/hammer-code/lms-be/domain"
	"github.com/sirupsen/logrus"
)

func (repo *repository) GetImage(ctx context.Context, fileName string) (data domain.Image, err error) {
	err = repo.db.DB(ctx).Where("file_name = ?", fileName).First(&data).Error
	if err != nil {
		logrus.Error("repo.Get Image: failed to get image")
		return
	}
	return data, nil
}
