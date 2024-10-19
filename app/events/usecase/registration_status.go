package usecase

import (
	"context"
	"errors"

	"github.com/hammer-code/lms-be/domain"
	"github.com/sirupsen/logrus"
)

func (uc usecase) RegistrationStatus(ctx context.Context, orderNo string) (resp domain.RegisterStatusResponse, err error) {
	rEvent, err := uc.repository.GetRegistrationEvent(ctx, orderNo)
	if err != nil {
		logrus.Error("failed to get event")
		return resp, err
	}

	if rEvent.ID == 0 {
		return resp, errors.New("registration order not found")
	}

	return domain.RegisterStatusResponse{
		OrderNo: orderNo,
		Status:  rEvent.Status,
	}, err
}
