package repository

import (
	"github.com/hammer-code/lms-be/app/users"
	"gorm.io/gorm"
)

type (
	repository struct {
		db *gorm.DB
	}
)

func NewRepository(db *gorm.DB) users.UserRepository {
	return &repository{
		db,
	}
}
