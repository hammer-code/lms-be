package http

import (
	"github.com/hammer-code/lms-be/domain"
)

type Handler struct {
	usecase domain.UserUsecase
}

var (
	handlr *Handler
)

func NewHandler(userUsecase domain.UserUsecase) domain.UserHandler {
	if handlr == nil {
		handlr = &Handler{
			usecase: userUsecase,
		}
	}

	return *handlr
}
