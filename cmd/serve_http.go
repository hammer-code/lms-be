package cmd

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	muxHandlers "github.com/gorilla/handlers"
	"github.com/gorilla/mux"

	"github.com/hammer-code/lms-be/app"
	"github.com/hammer-code/lms-be/config"
	"github.com/hammer-code/lms-be/domain"
	"github.com/hammer-code/lms-be/utils"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	httpSwagger "github.com/swaggo/http-swagger"
	"github.com/swaggo/swag"
)

var serveHttpCmd = &cobra.Command{
	Use:   "http",
	Short: "launches an HTTP server",
	Long:  "the serveHttp command initiates an HTTP server",
	Run: func(cmd *cobra.Command, args []string) {
		// load add package serve http here
		ctx := context.Background()

		cfg := config.GetConfig()

		app := app.InitApp(cfg)

		// route
		router := registerHandler(app)

		// build cors
		muxCorsWithRouter := muxHandlers.CORS(
			muxHandlers.AllowCredentials(),
			muxHandlers.AllowedHeaders(cfg.CORS_ALLOWED_HEADERS),
			muxHandlers.AllowedMethods(cfg.CORS_ALLOWED_METHODS),
			muxHandlers.AllowedOrigins(cfg.CORS_ALLOWED_ORIGINS),
		)(router)

		srv := &http.Server{
			Addr:    cfg.APP_PORT,
			Handler: muxCorsWithRouter,
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

		logrus.Info(fmt.Sprintf("server started, running on port %s", cfg.APP_PORT))
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			logrus.Fatal("starting server failed", err)
		}
	},
}

// func LoadJSON(path string) string {
// 	jsonBytes, err := os.ReadFile(path)

// 	// jsonBytes, err := os.ReadFile("documentation/users.json")
// 	if err != nil {
// 		fmt.Println("Error reading JSON file:", err)
// 		return ""
// 	}
// 	return string(jsonBytes)
// }

func LoadJSON(path string) string {
	jsonBytes, err := os.ReadFile(path)

	// jsonBytes, err := os.ReadFile("documentation/users.json")
	if err != nil {
		fmt.Println("Error reading JSON file:", err)
		return ""
	}
	return string(jsonBytes)
}

func LoadSwagger() {
	userTemplate := LoadJSON("documentation/users.json")
	var UsersSwaggerInfo = &swag.Spec{
		InfoInstanceName: "swagger",
		SwaggerTemplate:  userTemplate,
	}
	swag.Register(UsersSwaggerInfo.InstanceName(), UsersSwaggerInfo)
}

func init() {
	LoadSwagger()
	rootCmd.AddCommand(serveHttpCmd)

}

func health(w http.ResponseWriter, _ *http.Request) {
	utils.Response(domain.HttpResponse{
		Code:    200,
		Message: "good",
		Data:    nil,
	}, w)
}

func registerHandler(app app.App) *mux.Router {

	router := mux.NewRouter()

	router.HandleFunc("/health", health)

	doc := router.PathPrefix("/user")
	doc.Handler(httpSwagger.WrapHandler)

	v1 := router.PathPrefix("/api/v1").Subrouter()
	v1.HandleFunc("/newsletters/subscribe", app.NewLetterHandler.Subscribe).Methods(http.MethodPost)

	protectedV1Route := v1.NewRoute().Subrouter()
	protectedV1Route.Use(app.Middleware.AuthMiddleware)

	v1.HandleFunc("/register", app.UserHandler.Register).Methods(http.MethodPost)
	v1.HandleFunc("/login", app.UserHandler.Login).Methods(http.MethodPost)

	protectedV1Route.HandleFunc("/users", app.UserHandler.GetUsers).Methods(http.MethodGet)
	protectedV1Route.HandleFunc("/user", app.UserHandler.GetUserProfile).Methods(http.MethodGet)
	protectedV1Route.HandleFunc("/logout", app.UserHandler.Logout).Methods(http.MethodPost)

	protectedV1Route.HandleFunc("/", app.UserHandler.GetUserById).Methods(http.MethodGet)
	protectedV1Route.HandleFunc("/update", app.UserHandler.UpdateProfileUser).Methods(http.MethodPut)
	protectedV1Route.HandleFunc("/delete", app.UserHandler.DeleteUser).Methods(http.MethodDelete)

	return router
}
