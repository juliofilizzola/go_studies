package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Println("request feita")
		writer.Write([]byte("Hello World"))
	})
	fmt.Println("Server started at port 8080")
	log.Fatal(http.ListenAndServe(":5000", nil))
}
