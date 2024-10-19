package usecase

import (
	"context"

	"github.com/hammer-code/lms-be/domain"
)

func (uc usecase) GetEventByID(ctx context.Context, id uint) (resp domain.Event, err error) {
	return uc.repository.GetEvent(ctx, id)
}
