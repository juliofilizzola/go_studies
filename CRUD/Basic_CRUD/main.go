package main

import (
	"Basic_CRUD/services"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {
	fmt.Println("BASICO DE CRUD AQUI...")
	route := mux.NewRouter()
	route.HandleFunc("/user", services.CreateUser).Methods(http.MethodPost)
	fmt.Println("Server running...")
	log.Fatal(http.ListenAndServe(":5000", route))
}
