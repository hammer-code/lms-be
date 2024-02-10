package servehttp

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/hammer-code/lms-be/app/servehttp/handler"
	user_repository "github.com/hammer-code/lms-be/app/users/repository"
	user_usecase "github.com/hammer-code/lms-be/app/users/usecase"
	"github.com/hammer-code/lms-be/config"
	"github.com/hammer-code/lms-be/domain"
	"github.com/hammer-code/lms-be/utils"
	"gorm.io/gorm"
)

type (
	ContainerHTTP struct {
		Config   config.Config
		Database *gorm.DB
	}
)

// all route here
func Health(w http.ResponseWriter, r *http.Request) {
	utils.Response(domain.HttpResponse{
		Code:    200,
		Message: "good",
		Data:    nil,
	}, w)
}

func GetHandler(c ContainerHTTP) *mux.Router {
	// repository
	userRepository := user_repository.NewRepository(c.Database)
	// usecase
	userUsecase := user_usecase.NewUsecase(userRepository)
	h := handler.NewHandler()
	h.SetUsecase(userUsecase)
	h.Validate()

	r := mux.NewRouter()
	r.HandleFunc("/health", Health)

	r.HandleFunc("/users", h.GetUsers)

	return r
}
