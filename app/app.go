package app

import (
	"github.com/hammer-code/lms-be/app/middlewares"
	newsletters "github.com/hammer-code/lms-be/app/newsletters"
	users "github.com/hammer-code/lms-be/app/users"
	"github.com/hammer-code/lms-be/config"
	"github.com/hammer-code/lms-be/domain"
	pkgDB "github.com/hammer-code/lms-be/pkg/db"
	"github.com/hammer-code/lms-be/pkg/jwt"
	"gorm.io/driver/postgres"
)

type App struct {
	Middleware       domain.Middleware
	UserHandler      domain.UserHandler
	NewLetterHandler domain.NewslettterHandler
}

func InitApp(
	cfg config.Config,
) App {

	db := config.GetDatabase(postgres.Dialector{
		Config: &postgres.Config{
			DSN: cfg.DB_POSTGRES_DSN,
		}})

	dbTx := pkgDB.NewDBTransaction(db)
	jwtInstance := jwt.NewJwt(cfg.JWT_SECRET_KEY)

	// repository
	userRepo := users.InitRepository(dbTx)

	// Middlewares
	middleware := middlewares.InitMiddleware(jwtInstance, userRepo)

	userHandler := users.InitHandler(dbTx, jwtInstance)

	// repository
	newsletterRepo := newsletters.InitRepository(dbTx)
	// usecase
	newsletterUC := newsletters.InitUsecase(cfg, newsletterRepo, dbTx, jwt.NewJwt(cfg.JWT_SECRET_KEY))

	// handler
	newsletterHandler := newsletters.InitHandler(newsletterUC, middleware)

	return App{
		UserHandler:      userHandler,
		NewLetterHandler: newsletterHandler,
		Middleware:       middleware,
	}
}
