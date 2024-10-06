package usecase

import (
	"context"
	"errors"

	"github.com/hammer-code/lms-be/domain"
	"github.com/sirupsen/logrus"
)

func (uc usecase) CreatePayEvent(ctx context.Context, payload domain.EventPayPayload) error {
	rEvent, err := uc.repository.GetRegistrationEvent(ctx, payload.OrderNo)
	if err != nil {
		logrus.Error("failed to get event")
		return err
	}

	if rEvent.ID == 0 {
		return errors.New("registration order not found")
	}

	err = uc.dbTX.StartTransaction(ctx, func(txCtx context.Context) error {
		_, err := uc.repository.CreatePayEvent(txCtx, domain.EventPay{
			RegistrationEvent: rEvent.ID,
			ImageProofPayment: payload.ImageProofPayment,
			NetAmount:         payload.NetAmount,
		})
		if err != nil {
			logrus.Error("failed to get event")
			return err
		}
		return nil
	})

	return err
}
