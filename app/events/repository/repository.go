package repository

import (
	"context"

	"github.com/hammer-code/lms-be/domain"
	pkgDB "github.com/hammer-code/lms-be/pkg/db"
	"github.com/sirupsen/logrus"
)

type (
	repository struct {
		db pkgDB.DatabaseTransaction
	}
)

var (
	repo *repository
)

func NewRepository(db pkgDB.DatabaseTransaction) domain.EventRepository {
	if repo == nil {
		repo = &repository{
			db,
		}
	}

	return repo
}

func (r repository) CreateEvent(ctx context.Context, event domain.Event) (int, error) {
	err := repo.db.DB(ctx).Create(&event).Error
	if err != nil {
		logrus.Error("repo.CreateUser : failed to create user")
		return 0, err
	}
	return 1, nil
}

// 	UpdateEvent(ctx context.Context, payload UpdateEvenPayload) error
// 	DeleteEvent(ctx context.Context, id int) error
// 	GetEventByID(ctx context.Context, id int) (EventDTO, error)
// 	GetEvents(ctx context.Context, filter EventFilter) ([]EventDTO, error)
