package auth

import (
	"time"

	"github.com/golang-jwt/jwt/v5"

	"task-api/internal/config"
)

func GenerateToken(cfg *config.Config) (string, error) {

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user": "admin",
		"exp":  time.Now().Add(24 * time.Hour).Unix(),
	})

	return token.SignedString([]byte(cfg.JWT.Secret))
}
