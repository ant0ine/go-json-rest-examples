package main

import (
	"github.com/ant0ine/go-json-rest/rest"
	"log"
	"net/http"
)

func main() {
	router, err := rest.MakeRouter(
		&rest.Route{"GET", "/message.txt", func(w rest.ResponseWriter, req *rest.Request) {
			w.Header().Set("Content-Type", "text/plain")
			w.(http.ResponseWriter).Write([]byte("Hello World!"))
		}},
	)
	if err != nil {
		log.Fatal(err)
	}

	api := rest.NewApi(router)
	api.Use(rest.DefaultDevStack...)
	log.Fatal(http.ListenAndServe(":8080", api.MakeHandler()))
}
