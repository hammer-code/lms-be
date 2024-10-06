package usecase

import (
	"context"
)

func (us *usecase) UpdateUseImage(ctx context.Context, id uint) error {
	return us.imageRepo.UpdateUseImage(ctx, id)
}
