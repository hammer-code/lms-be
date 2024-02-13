package usecase

import (
	"github.com/hammer-code/lms-be/app/users"
	"github.com/hammer-code/lms-be/pkg/db"
)

type usecase struct {
	userRepo users.UserRepository
	dbTX     db.DatabaseTransaction
}

func NewUsecase(
	userRepo users.UserRepository,
	dbTX db.DatabaseTransaction,
) users.UserUsecase {
	return &usecase{
		userRepo: userRepo,
		dbTX:     dbTX,
	}
}
