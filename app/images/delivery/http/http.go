package http

import (
	"github.com/hammer-code/lms-be/domain"
)

type Handler struct {
	usecase domain.ImageUsecase
}

var (
	handlr *Handler
)

func NewHandler(usecase domain.ImageUsecase) domain.ImageHandler {
	if handlr == nil {
		handlr = &Handler{
			usecase: usecase,
		}
	}

	return *handlr
}
