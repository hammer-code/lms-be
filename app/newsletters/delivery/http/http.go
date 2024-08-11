package http

import (
	"github.com/hammer-code/lms-be/app/middlewares"
	"github.com/hammer-code/lms-be/app/newsletters"
)

type Handler struct {
	usecase    newsletters.NewslettersUsecase
	Middleware middlewares.Middleware
}

func NewHandler(ucNewsletter newsletters.NewslettersUsecase, middleware *middlewares.Middleware) Handler {
	return Handler{
		usecase:    ucNewsletter,
		Middleware: *middleware,
	}
}
