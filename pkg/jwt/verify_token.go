package jwt

import (
	"github.com/golang-jwt/jwt/v5"
)

func (j *jwtConfig) VerifyToken(token string) (*JwtCustomClaims, error) {
	tkn, err := jwt.ParseWithClaims(token, &JwtCustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(j.SecretKey), nil
	})

	if err != nil {
		return nil, err
	}

	return tkn.Claims.(*JwtCustomClaims), nil
}
