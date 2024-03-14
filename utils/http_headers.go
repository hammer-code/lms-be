package utils

import (
	"net/http"
	"strings"
)

func ExtractBearerToken(r *http.Request) *string {
	token := r.Header.Get("Authorization")
	if !strings.HasPrefix(token, "Bearer ") {
		return &token
	}

	token = strings.Split(token, "Bearer ")[1]
	return &token
}
