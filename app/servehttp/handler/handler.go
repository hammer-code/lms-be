package handler

import (
	"github.com/hammer-code/lms-be/app/users"
	"github.com/sirupsen/logrus"
)

type Handler struct {
	UserUsecase users.UserUsecase
}

func NewHandler() Handler {
	return Handler{}
}

func (h *Handler) SetUsecase(userUsecase users.UserUsecase) {
	h.UserUsecase = userUsecase
}

func (h *Handler) Validate() {
	if h.UserUsecase == nil {
		logrus.Fatal("handler: user usecase is nil")
	}
}
