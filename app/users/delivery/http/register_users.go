package http

import (
	"encoding/json"
	"github.com/hammer-code/lms-be/domain"
	"github.com/hammer-code/lms-be/utils"
	"github.com/sirupsen/logrus"
	"io"
	"net/http"
)

func (h Handler) Register(w http.ResponseWriter, r *http.Request) {
	bodyBytes, err := io.ReadAll(r.Body)
	if err != nil {
		utils.Response(domain.HttpResponse{
			Code:    500,
			Message: err.Error(),
		}, w)
		return
	}

	user := domain.Register{}
	if err = json.Unmarshal(bodyBytes, &user); err != nil {
		utils.Response(domain.HttpResponse{
			Code:    500,
			Message: err.Error(),
		}, w)
		return
	}

	if user.Password != user.ConfirmPassword {
		utils.Response(domain.HttpResponse{
			Code:    400,
			Message: "Confirm password doesnt match",
			Data:    nil,
		}, w)
		return
	}

	userInput := domain.RegistToUser(user)
	resultUser, err := h.usecase.Register(r.Context(), userInput)
	if err != nil {
		logrus.Error("userUsecase: failed to register user")
		utils.Response(domain.HttpResponse{
			Code:    500,
			Message: err.Error(),
		}, w)
		return
	}
	utils.Response(domain.HttpResponse{
		Code:    200,
		Message: "success",
		Data:    resultUser,
	}, w)
}
