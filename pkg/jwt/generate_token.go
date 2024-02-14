package jwt

import (
	"context"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/hammer-code/lms-be/domain"
)

func (j *jwtConfig) GenerateAccessToken(c context.Context, user *domain.User) (*string, error) {
	expiredTime := time.Now().Local().Add(60 * time.Minute)

	claims := JwtCustomClaims{
		ID:       user.ID,
		UserName: user.Username,
		Email:    user.Email,
		Role:     user.Role,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expiredTime),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	if encodedToken, err := token.SignedString([]byte(j.SecretKey)); err != nil {
		return nil, err
	} else {
		return &encodedToken, err
	}
}
