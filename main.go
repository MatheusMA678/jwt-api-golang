package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"matheus/jwt-api/auth"
	"matheus/jwt-api/handlers"

	"github.com/dgrijalva/jwt-go"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

// Structs
type HomeResponse struct {
	Message string `json:"message"`
}

type GetJwtStruct struct {
	Token string `json:"token"`
}

var key []byte

func init() {
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Erro ao carregar o arquivo .env")
	}

	key = []byte(os.Getenv("SECRET_KEY"))
}

type Credential struct{
	Username string `json:"username"`
	Password string `json:"password"`
}

var userDb = map[string]string{
	"user1": "password123",
}


func login(w http.ResponseWriter, r *http.Request) {
	// create a Credentials object
	var creds Credential
	// decode json to struct
	err := json.NewDecoder(r.Body).Decode(&creds)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	// verify if user exist or not
	userPassword, ok := userDb[creds.Username]

	// if user exist, verify the password
	if !ok || userPassword != creds.Password {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	// Create a token object and add the Username and StandardClaims
	var tokenClaim = auth.Token {
		Username: creds.Username,
		StandardClaims: jwt.StandardClaims{
			// Enter expiration in milisecond
			ExpiresAt: time.Now().Add(10 * time.Minute).Unix(),
		},
	}

	// Create a new claim with HS256 algorithm and token claim
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, tokenClaim )

	tokenString, err := token.SignedString(key)

	if err != nil {
		log.Fatal(err)
	}
	json.NewEncoder(w).Encode(tokenString)
}

// Main
func main() {
	r := mux.NewRouter()

	r.HandleFunc("/login", login).Methods("POST")
	r.HandleFunc("/me", handlers.Dashboard).Methods("GET")

	fmt.Println("Iniciando o servidor na porta 8000...")
	log.Fatal(http.ListenAndServe(":8000", r))
}