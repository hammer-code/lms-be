package usecase

import (
	"github.com/hammer-code/lms-be/domain"
	"github.com/hammer-code/lms-be/pkg/db"
	"github.com/hammer-code/lms-be/pkg/jwt"
)

type usecase struct {
	userRepo domain.UserRepository
	dbTX     db.DatabaseTransaction
	jwt      jwt.JWT
}

var (
	usec *usecase
)

func NewUsecase(userRepo domain.UserRepository, dbTX db.DatabaseTransaction, jwt jwt.JWT) domain.UserUsecase {
	if usec == nil {
		usec = &usecase{
			userRepo: userRepo,
			dbTX:     dbTX,
			jwt:      jwt,
		}
	}
	return usec
}
