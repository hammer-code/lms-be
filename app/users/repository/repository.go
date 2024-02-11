package repository

import (
	"github.com/hammer-code/lms-be/app/users"
)

type (
	repository struct {
	}
)

func NewRepository() users.UserRepository {
	return &repository{}
}
