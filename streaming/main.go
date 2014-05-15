package main

import (
	"fmt"
	"github.com/ant0ine/go-json-rest/rest"
	"net/http"
	"time"
)

func main() {

	handler := rest.ResourceHandler{
		EnableRelaxedContentType: true,
		DisableJsonIndent:        true,
	}
	handler.SetRoutes(
		&rest.Route{"GET", "/stream", StreamThings},
	)
	http.ListenAndServe(":8080", &handler)
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
