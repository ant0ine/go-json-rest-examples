package main

import (
	"github.com/ant0ine/go-json-rest/rest"
	"net/http"
)

func main() {
	handler := rest.ResourceHandler{}
	handler.SetRoutes(
		&rest.Route{"GET", "/message.txt", func(w rest.ResponseWriter, req *rest.Request) {
			w.Header().Set("Content-Type", "text/plain")
			w.(http.ResponseWriter).Write([]byte("Hello World!"))
		}},
	)
	http.ListenAndServe(":8080", &handler)
}
