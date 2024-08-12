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

func NewRepository(db pkgDB.DatabaseTransaction) domain.UserRepository {
	if repo == nil {
		repo = &repository{
			db,
		}
	}

	return repo
}
