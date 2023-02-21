package main

import (
	route "HTTP_2/http/Request"
	"fmt"
	"log"
	"net/http"
)

func main() {
	route.Route()
	fmt.Println("Running server")
	log.Fatal(http.ListenAndServe(":5000", nil))
}
