package middlewares

import (
	"github.com/hammer-code/lms-be/domain"
	"github.com/hammer-code/lms-be/pkg/jwt"
)

type Middleware struct {
	Jwt      jwt.JWT
	UserRepo domain.UserRepository
}

func InitMiddleware(jwt jwt.JWT, userRepo domain.UserRepository) domain.Middleware {
	return &Middleware{
		Jwt:      jwt,
		UserRepo: userRepo,
	}
}
