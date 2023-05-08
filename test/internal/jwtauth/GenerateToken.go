package jwtauth

import (
	"github.com/golang-jwt/jwt/v4"
)

func GereteToken(id interface{}) (string, error) {

	secret := "JWT_SECRET_KEY"
	claims := jwt.MapClaims{
		"id": id,
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	t, err := token.SignedString([]byte(secret))
	if err != nil {
		return "", err
	}
	return t, nil
}
