package main

import (
	"fmt"
	"github.com/ant0ine/go-json-rest/rest"
	"github.com/stretchr/graceful"
	"log"
	"net/http"
	"time"
)

func main() {
	api := rest.NewApi()
	api.Use(rest.DefaultDevStack...)
	router, err := rest.MakeRouter(
		rest.Get("/message", func(w rest.ResponseWriter, req *rest.Request) {
			for cpt := 1; cpt <= 10; cpt++ {

				// wait 1 second
				time.Sleep(time.Duration(1) * time.Second)

				w.WriteJson(map[string]string{
					"Message": fmt.Sprintf("%d seconds", cpt),
				})
				w.(http.ResponseWriter).Write([]byte("\n"))

				// Flush the buffer to client
				w.(http.Flusher).Flush()
			}
		}),
	)
	if err != nil {
		log.Fatal(err)
	}
	api.SetApp(router)

	server := &graceful.Server{
		Timeout: 10 * time.Second,
		Server: &http.Server{
			Addr:    ":8080",
			Handler: api.MakeHandler(),
		},
	}

	log.Fatal(server.ListenAndServe())
}
