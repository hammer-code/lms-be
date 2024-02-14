package jwt

import (
	"errors"

	"github.com/golang-jwt/jwt/v5"
)

func (j *jwtConfig) VerifyToken(token string) (*JwtCustomClaims, error) {
	tkn, err := jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
		if !t.Valid {
			return nil, errors.New("sign in to proceed")
		}

		return t, nil
	})

	if err != nil {
		return nil, err
	}

	return tkn.Claims.(*JwtCustomClaims), nil
}

