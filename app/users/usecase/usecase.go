package usecase

import (
	"github.com/hammer-code/lms-be/app/users"
)

type (
	usecase struct {
		repository users.UserRepository
	}
)

func NewUsecase(
	repo users.UserRepository,
) users.UserUsecase {
	return &usecase{
		repository: repo,
	}
}
