package routes

import (
	"fmt"
	"net/http"
)

func User(writer http.ResponseWriter, _request *http.Request) {
	fmt.Println("request feita")
	writer.Write([]byte("User 1"))
}
