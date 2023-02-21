package route

import "net/http"

func Route() {
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {

	})
}
