package repository

import (
	"github.com/hammer-code/lms-be/app/newsletters"
	pkgDB "github.com/hammer-code/lms-be/pkg/db"
)

type (
	repository struct {
		db pkgDB.DatabaseTransaction
	}
)

// injek magang
func NewRepository(db pkgDB.DatabaseTransaction) newsletters.NewsletterRepository {
	return &repository{
		db,
	}
}
