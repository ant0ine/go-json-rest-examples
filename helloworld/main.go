package main

import (
	"github.com/ant0ine/go-json-rest/rest"
	"log"
	"net/http"
)

type Message struct {
	Body string
}

func main() {
	router, err := rest.MakeRouter(
		&rest.Route{"GET", "/message", func(w rest.ResponseWriter, req *rest.Request) {
			w.WriteJson(&Message{
				Body: "Hello World!",
			})
		}},
	)
	if err != nil {
		log.Fatal(err)
	}

	api := rest.NewApi(router)
	api.Use(rest.DefaultDevStack...)
	log.Fatal(http.ListenAndServe(":8080", api.MakeHandler()))
}
