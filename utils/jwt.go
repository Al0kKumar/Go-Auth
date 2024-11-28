package utils

import (
	"time"
	"github.com/golang-jwt/jwt/v4"
)

var jwtkey = []byte("secret_key")

func GenerateJWT(email string) (string , error){
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email": email,
		"exp": time.Now().Add(24*time.Hour).Unix(),
	})

	return token.SignedString(jwtkey)
}