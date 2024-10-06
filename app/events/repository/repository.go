package repository

import (
	"github.com/hammer-code/lms-be/domain"
	pkgDB "github.com/hammer-code/lms-be/pkg/db"
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

// 	UpdateEvent(ctx context.Context, payload UpdateEvenPayload) error
// 	DeleteEvent(ctx context.Context, id int) error
// 	GetEventByID(ctx context.Context, id int) (EventDTO, error)
// 	GetEvents(ctx context.Context, filter EventFilter) ([]EventDTO, error)
