/* Demonstrate a streaming REST API, where the data is "flushed" to the client ASAP.

The Curl Demo:

        curl -i http://127.0.0.1:8080/stream

*/
package main

import (
	"fmt"
	"github.com/ant0ine/go-json-rest"
	"net/http"
	"time"
)

func main() {

	handler := rest.ResourceHandler{
		EnableRelaxedContentType: true,
	}
	handler.SetRoutes(
		rest.Route{"GET", "/stream", StreamThings},
	)
	http.ListenAndServe(":8080", &handler)
}

type Thing struct {
	Name string
}

func StreamThings(w *rest.ResponseWriter, r *rest.Request) {
	cpt := 0
	for {
		cpt++
		w.WriteJson(
			&Thing{
				Name: fmt.Sprintf("thing #%d", cpt),
			},
		)
		// Flush the buffer to client
		w.Flush()
		// wait 3 seconds
		time.Sleep(time.Duration(3) * time.Second)
	}
}
