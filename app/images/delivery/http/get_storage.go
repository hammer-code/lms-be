package http

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/hammer-code/lms-be/domain"
	"github.com/hammer-code/lms-be/utils"
	"github.com/sirupsen/logrus"
)

func (h Handler) GetStorage(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	// kind := vars["kind"]
	path := vars["path"]

	filePath, err := h.usecase.GetStorage(r.Context(), path)
	if err != nil {
		logrus.Error("failed to Create pay event : ", err)
		utils.Response(domain.HttpResponse{
			Code:    500,
			Message: err.Error(),
		}, w)
		return
	}

	if filePath == "" {
		logrus.Error("file not found")
		utils.Response(domain.HttpResponse{
			Code:    404,
			Message: "file not found",
		}, w)
		return
	}

	http.ServeFile(w, r, filePath)
}
