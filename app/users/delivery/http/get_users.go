package http

import (
	"encoding/json"
	"io"
	"net/http"

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

	userInput := domain.RegistToUser(user)
	resultUser, err := h.usecase.Register(r.Context(), userInput)
	if err != nil {
		logrus.Error("userUsecase: failed to regist user")
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

func (h Handler) Login(w http.ResponseWriter, r *http.Request) {
	bodyBytes, err := io.ReadAll(r.Body)
	if err != nil {
		utils.Response(domain.HttpResponse{
			Code:    500,
			Message: err.Error(),
		}, w)
		return
	}

	lmaoKang := domain.Login{}
	if err := json.Unmarshal(bodyBytes, &lmaoKang); err != nil {
		utils.Response(domain.HttpResponse{
			Code:    500,
			Message: err.Error(),
		}, w)
		return
	}

	_, token, err := h.usecase.Login(r.Context(), lmaoKang)

	if err != nil {
		utils.Response(domain.HttpResponse{
			Code:    500,
			Message: err.Error(),
		}, w)
		return
	}

	utils.Response(domain.HttpResponse{
		Code:    200,
		Message: "success",
		Data:    token,
	}, w)
}
