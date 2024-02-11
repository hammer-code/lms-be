package cmd

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/gorilla/mux"
	"github.com/hammer-code/lms-be/app/admins"
	admins_http "github.com/hammer-code/lms-be/app/admins/delivery/http"
	admins_usecase "github.com/hammer-code/lms-be/app/admins/usecase"
	users_repo "github.com/hammer-code/lms-be/app/users/repository"
	"github.com/hammer-code/lms-be/config"
	"github.com/hammer-code/lms-be/domain"
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

		// repository
		userRepo := users_repo.NewRepository()

		// usecase
		adminUsecase := admins_usecase.NewUsecase(db, userRepo)

		srv := &http.Server{
			Addr:    port(),
			Handler: registerHandler(adminUsecase),
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

func health(w http.ResponseWriter, r *http.Request) {
	utils.Response(domain.HttpResponse{
		Code:    200,
		Message: "good",
		Data:    nil,
	}, w)
}

func registerHandler(adminUsecase admins.UsecaseInterface) *mux.Router {

	router := mux.NewRouter()
	router.HandleFunc("/health", health)

	v1 := router.PathPrefix("/api/v1").Subrouter()

	admins_http.RegisterHandler(v1, adminUsecase)

	return router
}
