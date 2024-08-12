package domain

import "net/http"

type Middleware interface {
	AuthMiddleware(next http.Handler) http.Handler
}
