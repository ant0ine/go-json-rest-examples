package main

import (
	"github.com/ant0ine/go-json-rest-middleware-statsd"
	"github.com/ant0ine/go-json-rest/rest"
	"log"
	"net/http"
	"time"
)

func main() {
	api := rest.NewApi()
	api.Use(&statsd.StatsdMiddleware{})
	api.Use(rest.DefaultDevStack...)
	api.SetApp(rest.AppSimple(func(w rest.ResponseWriter, req *rest.Request) {

		// take more than 1ms so statsd can report it
		time.Sleep(100 * time.Millisecond)

		w.WriteJson(map[string]string{"Body": "Hello World!"})
	}))
	log.Fatal(http.ListenAndServe(":8080", api.MakeHandler()))
}
