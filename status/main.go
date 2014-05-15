package main

import (
	"github.com/ant0ine/go-json-rest/rest"
	"net/http"
)

func main() {
	handler := rest.ResourceHandler{
                EnableStatusService: true,
        }
	handler.SetRoutes(
		&rest.Route{"GET", "/.status",
			func(w rest.ResponseWriter, r *rest.Request) {
				w.WriteJson(handler.GetStatus())
			},
                },
	)
	http.ListenAndServe(":8080", &handler)
}
