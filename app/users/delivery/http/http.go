package http

import (
	"github.com/hammer-code/lms-be/app/middlewares"
	"github.com/hammer-code/lms-be/app/users"
)

type Handler struct {
	usecase    users.UserUsecase
	Middleware middlewares.Middleware
}

func NewHandler(userUsecase users.UserUsecase, middleware *middlewares.Middleware) Handler {
	return Handler{
		usecase:    userUsecase,
		Middleware: *middleware,
	}
}
