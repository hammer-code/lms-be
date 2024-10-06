package usecase

import (
	"context"

	"github.com/hammer-code/lms-be/domain"
	"github.com/sirupsen/logrus"
)

func (uc usecase) GetEvents(ctx context.Context, filter domain.EventFilter) (resp []domain.Event, pagination domain.Pagination, err error) {
	tData, resp, err := uc.repository.GetEvents(ctx, filter)
	if err != nil {
		logrus.Error("failed to get event")
		return
	}

	return resp, domain.NewPagination(tData, filter.FilterPagination), err
}
