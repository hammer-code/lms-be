package repository

import (
	"github.com/hammer-code/lms-be/app/users"
	pkgDB "github.com/hammer-code/lms-be/pkg/db"
)

type (
	repository struct {
		db pkgDB.DatabaseTransaction
	}
)

func NewRepository(db pkgDB.DatabaseTransaction) users.UserRepository {
	return &repository{
		db,
	}
}
