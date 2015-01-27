package main

import (
	"github.com/ant0ine/go-json-rest/rest"
	"log"
	"net/http"
)

func main() {
	statusMw := &rest.StatusMiddleware{}
	router, err := rest.MakeRouter(
		&rest.Route{"GET", "/.status",
			func(w rest.ResponseWriter, r *rest.Request) {
				w.WriteJson(statusMw.GetStatus())
			},
		},
	)
	if err != nil {
		log.Fatal(err)
	}

	api := rest.NewApi(router)
	api.Use(statusMw)
	api.Use(rest.DefaultDevStack...)
	log.Fatal(http.ListenAndServe(":8080", api.MakeHandler()))
}
