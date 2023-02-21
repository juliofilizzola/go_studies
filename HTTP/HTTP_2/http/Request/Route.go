package route

import (
	"html/template"
	"net/http"
)

var templetes *template.Template

type user struct {
	Name  string
	Email string
}

func Route() {
	templetes = template.Must(template.ParseGlob("./HTML/*.html"))
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		templetes.ExecuteTemplate(writer, "index.html", nil)
	})
}

func Route2() {
	templetes = template.Must(template.ParseGlob("./HTML/*.html"))
	http.HandleFunc("/user", func(writer http.ResponseWriter, request *http.Request) {
		user := user{
			Name:  "julio",
			Email: "juliofilizzola@hotmail.com",
		}
		templetes.ExecuteTemplate(writer, "index2.html", user)
	})
}
