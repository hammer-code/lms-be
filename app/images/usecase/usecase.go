package usecase

import (
	"github.com/hammer-code/lms-be/domain"
	"github.com/hammer-code/lms-be/pkg/db"
)

type usecase struct {
	imageRepo domain.ImageRepository
	dbTX      db.DatabaseTransaction
}

var (
	usec *usecase
)

func NewUsecase(imageRepo domain.ImageRepository, dbTX db.DatabaseTransaction) domain.ImageUsecase {
	if usec == nil {
		usec = &usecase{
			imageRepo: imageRepo,
			dbTX:      dbTX,
		}
	}
	return usec
}
