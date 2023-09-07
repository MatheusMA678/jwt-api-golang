package handlers

import (
	"encoding/json"
	"fmt"
	"matheus/jwt-api/auth"
	"net/http"

	"github.com/dgrijalva/jwt-go"
)

func Dashboard(w http.ResponseWriter, r *http.Request) {
	bearerToken := r.Header.Get("Authorization")

	token, err := auth.ValidateToken(bearerToken)

	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if !token.Valid {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	user := token.Claims.(*auth.Token)

	json.NewEncoder(w).Encode(fmt.Sprintf("%s Dashboard", user.Username))
}