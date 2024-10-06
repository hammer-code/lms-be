package http

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/hammer-code/lms-be/domain"
	"github.com/hammer-code/lms-be/utils"
	"github.com/sirupsen/logrus"
)

func (h Handler) RegisterEvent(w http.ResponseWriter, r *http.Request) {
	bodyBytes, err := io.ReadAll(r.Body)
	if err != nil {
		logrus.Error("failed to read body : ", err)
		utils.Response(domain.HttpResponse{
			Code:    500,
			Message: err.Error(),
		}, w)
		return
	}

	var payload domain.RegisterEventPayload
	if err := json.Unmarshal(bodyBytes, &payload); err != nil {
		logrus.Error("failed to unmarshal : ", err)
		utils.Response(domain.HttpResponse{
			Code:    500,
			Message: err.Error(),
		}, w)
		return
	}
	data, err := h.usecase.CreateRegisterEvent(r.Context(), payload)
	if err != nil {
		logrus.Error("failed to Create event : ", err)
		utils.Response(domain.HttpResponse{
			Code:    500,
			Message: err.Error(),
		}, w)
		return
	}

	utils.Response(domain.HttpResponse{
		Code:    201,
		Message: "success",
		Data:    data,
	}, w)
}
