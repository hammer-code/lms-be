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

// injek magang
func NewRepository(db pkgDB.DatabaseTransaction) domain.NewsletterRepository {
	return &repository{
		db,
	}
}
