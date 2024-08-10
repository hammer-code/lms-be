package app

import (
	"github.com/hammer-code/lms-be/app/middlewares"
	users "github.com/hammer-code/lms-be/app/users"
	"github.com/hammer-code/lms-be/config"
	"github.com/hammer-code/lms-be/domain"
	pkgDB "github.com/hammer-code/lms-be/pkg/db"
	"github.com/hammer-code/lms-be/pkg/jwt"
	"gorm.io/driver/postgres"
)

type App struct {
	UserHandler domain.UserHandler
	Middleware  domain.Middleware
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

	return App{
		UserHandler: userHandler,
		Middleware:  middleware,
	}
}
