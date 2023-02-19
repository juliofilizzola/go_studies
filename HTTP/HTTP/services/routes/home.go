package routes

import (
	"fmt"
	"net/http"
)

func Home(writer http.ResponseWriter, _request *http.Request) {
	fmt.Println("request feita")
	writer.Write([]byte("Hello World"))
}
