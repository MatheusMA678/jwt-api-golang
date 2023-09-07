package main

import (
	"fmt"
	"log"
	"net/http"

	"matheus/jwt-api/handlers"

	"github.com/gorilla/mux"
)

// Structs
type HomeResponse struct {
	Message string `json:"message"`
}

type GetJwtStruct struct {
	Token string `json:"token"`
}

// Main
func main() {
	r := mux.NewRouter()

	r.HandleFunc("/login", handlers.Login)
	r.HandleFunc("/me", handlers.Dashboard)

	fmt.Println("Iniciando o servidor na porta 8000...")
	log.Fatal(http.ListenAndServe(":8000", r))
}