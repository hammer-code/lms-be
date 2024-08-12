package newsletters

import (
	"github.com/hammer-code/lms-be/config"
	"github.com/hammer-code/lms-be/domain"
	"github.com/hammer-code/lms-be/pkg/db"
	"github.com/hammer-code/lms-be/pkg/jwt"

	newsletter_handler "github.com/hammer-code/lms-be/app/newsletters/delivery/http"
	newsletter_repo "github.com/hammer-code/lms-be/app/newsletters/repository"
	newsletter_usecase "github.com/hammer-code/lms-be/app/newsletters/usecase"
)

func InitRepository(dbTx db.DatabaseTransaction) domain.NewsletterRepository {
	return newsletter_repo.NewRepository(dbTx)
}

func InitUsecase(cfg config.Config, newsletterRepo domain.NewsletterRepository, dbTx db.DatabaseTransaction, jwt jwt.JWT) domain.NewslettersUsecase {
	return newsletter_usecase.NewUsecase(cfg, newsletterRepo, dbTx, jwt)
}

func InitHandler(newsletterUC domain.NewslettersUsecase, middleware domain.Middleware) domain.NewslettterHandler {
	return newsletter_handler.NewHandler(newsletterUC, middleware)
}
