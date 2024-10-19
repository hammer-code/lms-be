package usecase

import (
	"context"
	"os"
)

func (us *usecase) GetStorage(ctx context.Context, fileName string) (filePath string, err error) {
	im, err := us.imageRepo.GetImage(ctx, fileName)
	if err != nil {
		return "", err
	}

	if _, err := os.Stat(im.FilePath); os.IsNotExist(err) {
		return "", err
	}

	return im.FilePath, err
}
