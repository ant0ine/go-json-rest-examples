package main

import (
	"github.com/ant0ine/go-json-rest/rest"
	"net/http"
)

type Message struct {
	Body string
}

func main() {
	handler := rest.ResourceHandler{}
	handler.SetRoutes(
		&rest.Route{"GET", "/message", func(w rest.ResponseWriter, req *rest.Request) {
			w.WriteJson(&Message{
				Body: "Hello World!",
			})
		}},
	)
	http.Handle("/api/", http.StripPrefix("/api", &handler))

	http.Handle("/static/", http.StripPrefix("/static", http.FileServer(http.Dir("."))))

	http.ListenAndServe(":8080", nil)
}
