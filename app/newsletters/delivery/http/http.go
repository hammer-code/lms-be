package http

import (
	"github.com/hammer-code/lms-be/domain"
)

type handler struct {
	usecase    domain.NewslettersUsecase
	middleware domain.Middleware
}

var (
	handlr *handler
)

func NewHandler(ucNewsletter domain.NewslettersUsecase, middleware domain.Middleware) domain.NewslettterHandler {

	if handlr == nil {
		handlr = &handler{
			usecase:    ucNewsletter,
			middleware: middleware,
		}
	}
	return handlr
}
