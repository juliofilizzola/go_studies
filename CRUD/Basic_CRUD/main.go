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
	route.HandleFunc("/user", services.GetUsers).Methods(http.MethodGet)
	route.HandleFunc("/user/{id}", services.GetUser).Methods(http.MethodGet)
	route.HandleFunc("/user/{id}", services.UpdateUser).Methods(http.MethodPut)
	route.HandleFunc("/user/{id}", services.DeleteUser).Methods(http.MethodDelete)
	fmt.Println("Server running...")
	log.Fatal(http.ListenAndServe(":5000", route))
}
