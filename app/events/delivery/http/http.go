package http

import (
	"github.com/hammer-code/lms-be/domain"
)

type Handler struct {
	usecase domain.EventUsecase
}

var (
	handlr *Handler
)

func NewHandler(usecase domain.EventUsecase) domain.EventHandler {
	if handlr == nil {
		handlr = &Handler{
			usecase: usecase,
		}
	}

	return *handlr
}
