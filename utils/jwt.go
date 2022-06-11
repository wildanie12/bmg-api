package utils

import (
	_domain "gin-bmg-restful/entities/domain"
	"time"

	"github.com/golang-jwt/jwt"
)

const secret = "thisissecret33"

// CreateToken jwt
func CreateToken(user _domain.User) (string, error) {
	claims := jwt.MapClaims{
		"username": user.Username,
		"email": user.Email,
		"exp": time.Now().Add(time.Hour * 48).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(secret))
}