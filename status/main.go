/* Demonstrate how to setup a /.status endpoint

The Curl Demo:

        curl -i http://127.0.0.1:8080/.status
        curl -i http://127.0.0.1:8080/.status
        ...

*/
package main

import (
	"github.com/ant0ine/go-json-rest"
	"net/http"
)

func main() {
	handler := rest.ResourceHandler{
                EnableStatusService: true,
        }
	handler.SetRoutes(
		rest.Route{"GET", "/.status",
			func(w rest.ResponseWriter, r *rest.Request) {
				w.WriteJson(handler.GetStatus())
			},
                },
	)
	http.ListenAndServe(":8080", &handler)
}
