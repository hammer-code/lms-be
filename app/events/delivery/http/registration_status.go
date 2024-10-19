package http

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/hammer-code/lms-be/domain"
	"github.com/hammer-code/lms-be/utils"
	"github.com/sirupsen/logrus"
)

func (h Handler) RegistrationStatus(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	order_no := vars["order_no"]
	data, err := h.usecase.RegistrationStatus(r.Context(), order_no)
	if err != nil {
		logrus.Error("failed to Create pay event : ", err)
		utils.Response(domain.HttpResponse{
			Code:    500,
			Message: err.Error(),
		}, w)
		return
	}

	utils.Response(domain.HttpResponse{
		Code:    200,
		Message: "success",
		Data:    data,
	}, w)
}
