package usecase

import (
	"github.com/hammer-code/lms-be/app/users"
)

type usecase struct {
	userRepo users.UserRepository
}

func NewUsecase(
	userRepo users.UserRepository,
) users.UserUsecase {
	return &usecase{
		userRepo: userRepo,
	}
}
