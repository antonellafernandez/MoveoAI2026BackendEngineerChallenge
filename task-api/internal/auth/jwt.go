package auth

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var JwtSecret = []byte("my-secret-key")

func GenerateToken() (string, error) {

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user": "admin",
		"exp":  time.Now().Add(24 * time.Hour).Unix(),
	})

	return token.SignedString(JwtSecret)
}
