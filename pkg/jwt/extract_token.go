package jwt

import (
	"fmt"
	"strconv"

	"github.com/golang-jwt/jwt/v5"
)

// func ExtractToken(tokenString, secretKey string) (int, error) {
// 	if len(tokenString) >= 7 && tokenString[0:7] == "Bearer " {
// 		tokenString = tokenString[7:]
// 	}

// 	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
// 		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
// 			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
// 		}
// 		return []byte(secretKey), nil
// 	}); 
// 	if err != nil {
// 		return 0, err
// 	}

// 	claims := token.Claims.(jwt.MapClaims)
// 	id := claims["ID"].(string)
// 	idInt,_:= strconv.Atoi(id)

// 	return idInt, nil

// }

func  ParseToken(tokenString, secretKey string) (*JwtCustomClaims, error) {
	// Parse token
	token, err := jwt.ParseWithClaims(tokenString, &JwtCustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(secretKey), nil
	})
	if err != nil {
		return nil, err
	}

	// Check token validity
	if claims, ok := token.Claims.(*JwtCustomClaims); ok && token.Valid {
		return claims, nil
	} else {
		return nil, errors.New("invalid token")
	}
}


func

