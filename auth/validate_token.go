package auth

import (
	"strings"

	"github.com/dgrijalva/jwt-go"
)

type Token struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

func ValidateToken(bearerToken string) (*jwt.Token, error) {

	// format the token string
	tokenStr := strings.Split(bearerToken, " ")[1]

	// parse the tonken with tokenObj
	token, err := jwt.ParseWithClaims(tokenStr, &Token{}, func(t *jwt.Token) (interface{}, error) {
		return t, nil
	})

	// return token and error
	return token, err
}