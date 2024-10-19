package repository

import (
	"context"

	"github.com/hammer-code/lms-be/domain"
	"github.com/sirupsen/logrus"
)

// query untuk get user di database
func (repo *repository) GetEvents(ctx context.Context, filter domain.EventFilter) (tData int, data []domain.Event, err error) {
	db := repo.db.DB(ctx).Model(&domain.Event{})

	var totalData int64

	if filter.Type != "" {
		db = db.Where("type = ?", filter.Type)
	}

	if filter.Status != "" {
		db = db.Where("status = ?", filter.Status)
	}

	if filter.StartDate.Valid {
		db = db.Where("start_date > ?", filter.StartDate)
	}

	if filter.StartDate.Valid {
		db = db.Where("end_date < ?", filter.EndDate)
	}

	if filter.Title != "" {
		db = db.Where("title ILIKE ?", "%"+filter.Title+"%")
	}

	db.Count(&totalData)

	err = db.Limit(filter.FilterPagination.GetLimit()).
		Offset(filter.FilterPagination.GetOffset()).
		Preload("Tags").Preload("Speakers").Find(&data).Error
	if err != nil {
		logrus.Error("repo.GetEvents: failed to get events use generic conditions")
		return
	}

	return int(totalData), data, err
}

func (repo *repository) GetEvent(ctx context.Context, eventID uint) (data domain.Event, err error) {
	db := repo.db.DB(ctx).Model(&domain.Event{})

	err = db.Where("id = ?", eventID).Find(&data).Error
	if err != nil {
		logrus.Error("repo.GetEvents: failed to get events use generic conditions")
		return
	}

	return data, err
}

func (repo *repository) GetRegistrationEvent(ctx context.Context, orderNo string) (data domain.RegistrationEvent, err error) {
	db := repo.db.DB(ctx).Model(&domain.RegistrationEvent{})

	err = db.Where("order_no = ?", orderNo).Find(&data).Error
	if err != nil {
		logrus.Error("failed to get registration event use generic conditions")
		return
	}

	return data, err
}

func (repo *repository) ListRegistration(ctx context.Context, filter domain.EventFilter) (tData int, data []domain.RegistrationEvent, err error) {
	db := repo.db.DB(ctx).Model(&domain.RegistrationEvent{})

	var totalData int64

	if filter.Status != "" {
		db = db.Where("status = ?", filter.Status)
	}

	if filter.StartDate.Valid {
		db = db.Where("start_date > ?", filter.StartDate)
	}

	if filter.StartDate.Valid {
		db = db.Where("end_date < ?", filter.EndDate)
	}

	db.Count(&totalData)

	err = db.Limit(filter.FilterPagination.GetLimit()).
		Offset(filter.FilterPagination.GetOffset()).Find(&data).Error
	if err != nil {
		logrus.Error("failed to list registration event use generic conditions")
		return
	}

	return int(totalData), data, err
}
