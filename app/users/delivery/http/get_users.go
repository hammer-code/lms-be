package http

import (
	"net/http"
	"strconv"

	"github.com/hammer-code/lms-be/domain"
	"github.com/hammer-code/lms-be/utils"
	"github.com/sirupsen/logrus"
)

func (h Handler) GetUsers(w http.ResponseWriter, r *http.Request) {
	users, err := h.usecase.GetUsers(r.Context())

	if err != nil {
		logrus.Error("userUsecase: failed to get users")
		utils.Response(domain.HttpResponse{
			Code:    500,
			Message: err.Error(),
		}, w)
		return
	}

	utils.Response(domain.HttpResponse{
		Code:    200,
		Message: "success",
		Data:    users,
	}, w)
}

func (h Handler) GetUserById(w http.ResponseWriter, r *http.Request) {
	userIDStr := r.URL.Query().Get("id")
	userID, _ := strconv.ParseUint(userIDStr, 10, 64)
	user, err := h.usecase.GetUserById(r.Context(), int8(userID))

	if err != nil {
		logrus.Error("userUsecase: failed to get user")
		utils.Response(domain.HttpResponse{
			Code:    500,
			Message: err.Error(),
		}, w)
		return
	}

	utils.Response(domain.HttpResponse{
		Code:    200,
		Message: "success",
		Data:    user,
	}, w)
}
