package cmd

import (
	"context"
	"fmt"
	"github.com/hammer-code/lms-be/app/middlewares"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/gorilla/mux"
	muxHandlers "github.com/gorilla/handlers"
	users_handler "github.com/hammer-code/lms-be/app/users/delivery/http"
	users_repo "github.com/hammer-code/lms-be/app/users/repository"
	users_usecase "github.com/hammer-code/lms-be/app/users/usecase"
	"github.com/hammer-code/lms-be/config"
	"github.com/hammer-code/lms-be/domain"
	pkgDB "github.com/hammer-code/lms-be/pkg/db"
	"github.com/hammer-code/lms-be/pkg/jwt"
	"github.com/hammer-code/lms-be/utils"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"gorm.io/driver/postgres"
)

func port() string {
	return ":8000"
}

var serveHttpCmd = &cobra.Command{
	Use:   "http",
	Short: "launches an HTTP server",
	Long:  "the serveHttp command initiates an HTTP server",
	Run: func(cmd *cobra.Command, args []string) {
		// load add package serve http here
		ctx := context.Background()

		cfg := config.GetConfig()

		db := config.GetDatabase(postgres.Dialector{
			Config: &postgres.Config{
				DSN: cfg.DB_POSTGRES_DSN,
			}})

		dbTx := pkgDB.NewDBTransaction(db)
		jwtInstance := jwt.NewJwt(cfg.JWT_SECRET_KEY)

		// repository
		userRepo := users_repo.NewRepository(dbTx)

		// usecase
		userUsecase := users_usecase.NewUsecase(userRepo, dbTx, jwt.NewJwt(cfg.JWT_SECRET_KEY))

		// Middlewares
		middleware := middlewares.Middleware{
			Jwt:      jwtInstance,
			UserRepo: userRepo,
		}

		// handler
		userHandler := users_handler.NewHandler(userUsecase, &middleware)

		srv := &http.Server{
			Addr: port(),
			Handler: registerHandler(handler{
				userHandler: userHandler,
			}),
		}

		go func() {
			done := make(chan os.Signal, 1)

			signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
			<-done
			logrus.Info("svr.Shutdown: start shutdown")
			if err := srv.Shutdown(ctx); err == context.DeadlineExceeded {
				logrus.Error("svr.Shutdown: context deadline exceeded", err)
			}
			logrus.Info("svr.Shutdown: shutdown success")
		}()

		logrus.Info(fmt.Sprintf("server started, running on port %s", port()))
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			logrus.Fatal("starting server failed", err)
		}
	},
}

func init() {
	rootCmd.AddCommand(serveHttpCmd)
}

func health(w http.ResponseWriter, _ *http.Request) {
	utils.Response(domain.HttpResponse{
		Code:    200,
		Message: "good",
		Data:    nil,
	}, w)
}

type handler struct {
	userHandler users_handler.Handler
}

func registerHandler(h handler) *mux.Router {

	router := mux.NewRouter()
	router.Use(muxHandlers.CORS())

	router.HandleFunc("/health", health)

	v1 := router.PathPrefix("/api/v1").Subrouter()

	protectedV1Route := v1.NewRoute().Subrouter()
	protectedV1Route.Use(h.userHandler.Middleware.AuthMiddleware)

	v1.HandleFunc("/register", h.userHandler.Register).Methods(http.MethodPost)
	v1.HandleFunc("/login", h.userHandler.Login).Methods(http.MethodPost)

	protectedV1Route.HandleFunc("/users", h.userHandler.GetUsers).Methods(http.MethodGet)
	protectedV1Route.HandleFunc("/logout", h.userHandler.Logout).Methods(http.MethodPost)

	return router
}
