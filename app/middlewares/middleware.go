package middlewares

import (
	"github.com/hammer-code/lms-be/app/users"
	"github.com/hammer-code/lms-be/pkg/jwt"
)

type Middleware struct {
	Jwt      jwt.JWT
	UserRepo users.UserRepository
}
