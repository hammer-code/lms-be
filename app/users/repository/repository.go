package repository

import (
	"github.com/hammer-code/lms-be/app/users"
	"gorm.io/gorm"
)

type (
	repository struct {
		database *gorm.DB
	}
)

func NewRepository(database *gorm.DB) users.UserRepository {
	return &repository{
		database: database,
	}
}
