package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {
	fmt.Println("BASICO DE CRUD AQUI...")
	route := mux.NewRouter()

	fmt.Println("Server running...")
	log.Fatal(http.ListenAndServe(":5000", route))
}
