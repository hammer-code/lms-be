package usecase

import (
	"github.com/hammer-code/lms-be/app/newsletters"
	"github.com/hammer-code/lms-be/config"
	"github.com/hammer-code/lms-be/pkg/db"
	"github.com/hammer-code/lms-be/pkg/jwt"
)

type usecase struct {
	cfg            config.Config
	newsLetterRepo newsletters.NewsletterRepository
	dbTX           db.DatabaseTransaction
	jwt            jwt.JWT
}

func NewUsecase(cfg config.Config, newsLetterRepo newsletters.NewsletterRepository, dbTX db.DatabaseTransaction, jwt jwt.JWT) newsletters.NewslettersUsecase {
	return &usecase{
		cfg:            cfg,
		newsLetterRepo: newsLetterRepo,
		dbTX:           dbTX,
		jwt:            jwt,
	}
}
