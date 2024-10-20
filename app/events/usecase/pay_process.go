package usecase

import (
	"context"
	"errors"

	"github.com/hammer-code/lms-be/domain"
	"github.com/sirupsen/logrus"
)

func (uc usecase) PayProcess(ctx context.Context, payload domain.PayProcessPayload) error {
	rEvent, err := uc.repository.GetRegistrationEvent(ctx, payload.OrderNo)
	if err != nil {
		logrus.Error("failed to get event")
		return err
	}

	if rEvent.ID == 0 {
		return errors.New("registration order not found")
	}

	if rEvent.Status == "SUCCESS" {
		return nil
	}

	eventPay, err := uc.repository.GetEventPay(ctx, payload.OrderNo)
	if err != nil {
		logrus.Error("failed to get event")
		return err
	}

	if eventPay.ID == 0 {
		return errors.New("event pay order not found")
	}

	if eventPay.Status == "SUCCESS" {
		return nil
	}

	eventPay.Status = payload.Status
	rEvent.Status = payload.Status
	rEvent.UpToYou = payload.Note

	err = uc.dbTX.StartTransaction(ctx, func(txCtx context.Context) error {
		err = uc.repository.UpdateEventPay(txCtx, eventPay)

		if err != nil {
			logrus.Error("failed to update event pay")
			return err
		}

		err = uc.repository.UpdateRegistrationEvent(txCtx, rEvent)
		if err != nil {
			logrus.Error("failed to update registration event pay")
			return err
		}

		return nil
	})

	return err
}
