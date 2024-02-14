package usecase

import (
	"github.com/hammer-code/lms-be/app/users"
	"github.com/hammer-code/lms-be/pkg/db"
	"github.com/hammer-code/lms-be/pkg/jwt"
)

type usecase struct {
	userRepo users.UserRepository
	dbTX     db.DatabaseTransaction
	jwt 	jwt.JWT
}

func NewUsecase( userRepo users.UserRepository, dbTX db.DatabaseTransaction, jwt jwt.JWT) users.UserUsecase {
	return &usecase{
		userRepo: userRepo,
		dbTX:     dbTX,
		jwt: 	  jwt,
	}
}
