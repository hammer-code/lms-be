package usecase

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/hammer-code/lms-be/domain"
	"github.com/hammer-code/lms-be/pkg/hash"
	"github.com/sirupsen/logrus"
)

func (uc usecase) CreateRegisterEvent(ctx context.Context, payload domain.RegisterEventPayload) (domain.RegisterEventResponse, error) {
	event, err := uc.repository.GetEvent(ctx, payload.EventID)
	if err != nil {
		logrus.Error("failed to get event")
		return domain.RegisterEventResponse{}, err
	}

	if event.ID == 0 {
		return domain.RegisterEventResponse{}, errors.New("event not found")
	}

	tNow := time.Now()

	if !event.BookingStart.Valid {
		return domain.RegisterEventResponse{}, errors.New("event is not start to booking")
	}

	if event.BookingEnd.Valid {
		if tNow.After(event.BookingEnd.Time) {
			return domain.RegisterEventResponse{}, errors.New("priode booking has ended")
		}
	}

	hash := hash.GenerateHash(time.Now().Format("2006-01-02 15:04:05"))

	orderNo := fmt.Sprintf("TXE-%d-%s%s%s%s", event.ID, time.Now().Format("06"), time.Now().Format("01"), time.Now().Format("02"), hash[0:4])

	err = uc.dbTX.StartTransaction(ctx, func(txCtx context.Context) error {
		_, err := uc.repository.CreateRegisterEvent(ctx, domain.RegistrationEvent{
			OrderNo:     orderNo,
			EventID:     event.ID,
			Name:        payload.Name,
			Email:       payload.Email,
			PhoneNumber: payload.PhoneNumber,
			Status:      "register",
			UpToYou:     "new register",
		})
		if err != nil {
			logrus.Error("failed to get event")
			return err
		}
		return nil
	})

	return domain.RegisterEventResponse{
		OrderNo: orderNo,
	}, err
}
