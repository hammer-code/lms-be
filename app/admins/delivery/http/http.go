package http

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/hammer-code/lms-be/app/admins"
)

type handler struct {
	usecase admins.UsecaseInterface
}

func RegisterHandler(v1 *mux.Router, adminUsecase admins.UsecaseInterface) {
	handler := handler{
		usecase: adminUsecase,
	}
	v1.HandleFunc("/users", handler.GetUsers).Methods(http.MethodGet)
}
