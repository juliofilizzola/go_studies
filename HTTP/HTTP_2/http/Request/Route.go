package route

import (
	"html/template"
	"net/http"
)

var templetes *template.Template

func Route() {
	templetes = template.Must(template.ParseGlob("./HTML/*.html"))
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		templetes.ExecuteTemplate(writer, "index.html", nil)
	})
}
