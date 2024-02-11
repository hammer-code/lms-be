package http

import (
	"github.com/hammer-code/lms-be/app/users"
)

type Handler struct {
	usecase users.UserUsecase
}

func NewHandler(userUsecase users.UserUsecase) Handler {
	return Handler{
		usecase: userUsecase,
	}
}
