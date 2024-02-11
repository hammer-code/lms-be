package usecase

import (
	"github.com/hammer-code/lms-be/app/admins"
	"github.com/hammer-code/lms-be/app/users"
	"gorm.io/gorm"
)

type usecase struct {
	db       *gorm.DB
	userRepo users.UserRepository
}

func NewUsecase(
	db *gorm.DB,
	userRepo users.UserRepository,
) admins.UsecaseInterface {
	return &usecase{
		db:       db,
		userRepo: userRepo,
	}
}
