package usecase

import (
	"context"

	"github.com/hammer-code/lms-be/domain"
	"github.com/sirupsen/logrus"
)

func (uc usecase) ListEventPay(ctx context.Context, filter domain.EventFilter) (resp []domain.EventPay, pagination domain.Pagination, err error) {
	tData, datas, err := uc.repository.ListEventPay(ctx, filter)
	if err != nil {
		logrus.Error("failed to list event pay")
		return
	}

	return datas, domain.NewPagination(tData, filter.FilterPagination), err
}
