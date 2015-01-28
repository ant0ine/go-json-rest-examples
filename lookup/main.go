package main

import (
	"github.com/ant0ine/go-json-rest/rest"
	"log"
	"net"
	"net/http"
)

type Message struct {
	Body string
}

func main() {
	api := rest.NewApi()
	api.Use(rest.DefaultDevStack...)
	router, err := rest.MakeRouter(
		&rest.Route{"GET", "/lookup/#host", func(w rest.ResponseWriter, req *rest.Request) {
			ip, err := net.LookupIP(req.PathParam("host"))
			if err != nil {
				rest.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			w.WriteJson(&ip)
		}},
	)
	if err != nil {
		log.Fatal(err)
	}
	api.SetApp(router)
	log.Fatal(http.ListenAndServe(":8080", api.MakeHandler()))
}
