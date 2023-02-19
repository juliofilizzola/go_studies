package main

import (
	"HTTP/services/routes"
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", routes.Home)
	http.HandleFunc("/user", routes.User)
	fmt.Println("Server started at port 8080")
	log.Fatal(http.ListenAndServe(":5000", nil))
}
