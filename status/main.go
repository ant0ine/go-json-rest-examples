package main

import (
	"github.com/ant0ine/go-json-rest/rest"
	"log"
	"net/http"
)

func main() {
	handler := rest.ResourceHandler{
		EnableStatusService: true,
	}
	err := handler.SetRoutes(
		&rest.Route{"GET", "/.status",
			func(w rest.ResponseWriter, r *rest.Request) {
				w.WriteJson(handler.GetStatus())
			},
		},
	)
	if err != nil {
		log.Fatal(err)
	}
	log.Fatal(http.ListenAndServe(":8080", &handler))
}
