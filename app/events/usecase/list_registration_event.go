package usecase

import (
	"context"

	"github.com/hammer-code/lms-be/domain"
	"github.com/sirupsen/logrus"
)

func (uc usecase) ListRegistration(ctx context.Context, filter domain.EventFilter) (resp []domain.RegistrationEvent, pagination domain.Pagination, err error) {
	tData, datas, err := uc.repository.ListRegistration(ctx, filter)
	if err != nil {
		logrus.Error("failed to get event")
		return
	}

	return datas, domain.NewPagination(tData, filter.FilterPagination), err
}
