package cmd

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/hammer-code/lms-be/app/servehttp"
	"github.com/hammer-code/lms-be/config"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
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

		maxConnIdel := viper.GetInt("DB_POSTGRES_MAX_CONN_IDLE")
		if maxConnIdel == 0 {
			maxConnIdel = 10
		}
		maxConnOpen := viper.GetInt("DB_POSTGRES_MAX_CONN_OPEN")
		if maxConnOpen == 0 {
			maxConnOpen = 10
		}

		handler := servehttp.GetHandler(servehttp.ContainerHTTP{
			Config: cfg,
			Database: config.GetDatabase(postgres.Dialector{
				Config: &postgres.Config{
					DSN: cfg.DB_POSTGRES_DSN,
				},
			}, maxConnIdel, maxConnOpen),
		})

		srv := &http.Server{
			Addr:    port(),
			Handler: handler,
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
