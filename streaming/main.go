package main

import (
	"fmt"
	"github.com/ant0ine/go-json-rest/rest"
	"log"
	"net/http"
	"time"
)

func main() {

	router, err := rest.MakeRouter(
		&rest.Route{"GET", "/stream", StreamThings},
	)
	if err != nil {
		log.Fatal(err)
	}

	api := rest.NewApi(router)
	api.Use(&rest.AccessLogApacheMiddleware{})
	api.Use(&rest.TimerMiddleware{})
	api.Use(&rest.RecorderMiddleware{})
	api.Use(&rest.RecoverMiddleware{})

	log.Fatal(http.ListenAndServe(":8080", api.MakeHandler()))
}

type Thing struct {
	Name string
}

func StreamThings(w rest.ResponseWriter, r *rest.Request) {
	cpt := 0
	for {
		cpt++
		w.WriteJson(
			&Thing{
				Name: fmt.Sprintf("thing #%d", cpt),
			},
		)
		w.(http.ResponseWriter).Write([]byte("\n"))
		// Flush the buffer to client
		w.(http.Flusher).Flush()
		// wait 3 seconds
		time.Sleep(time.Duration(3) * time.Second)
	}
}
