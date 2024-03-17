package jwt

import (
	"errors"

	"github.com/golang-jwt/jwt/v5"
)

func ParseToken(tokenString, secretKey string) (*JwtCustomClaims, error) {
	if len(tokenString) >= 7 && tokenString[0:7] == "Bearer " {
		tokenString = tokenString[7:]
	}
	
	token, err := jwt.ParseWithClaims(tokenString, &JwtCustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(secretKey), nil
	})
	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(*JwtCustomClaims); ok && token.Valid {
		return claims, nil
	} else {
		return nil, errors.New("invalid token")
	}
}
