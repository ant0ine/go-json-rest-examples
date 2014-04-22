/* Combine Go-Json-Rest with other handlers.

rest.ResourceHandler is a valid http.Handler, and can be combined with other handlers.
In this example the ResourceHandler is used under the /api/ prefix, while a FileServer
is instantiated under the /static/ prefix.

The curl demo:

        curl -i http://127.0.0.1:8080/api/message
        curl -i http://127.0.0.1:8080/static/main.go

*/
package main

import (
	"github.com/ant0ine/go-json-rest/rest"
	"net/http"
)

type Message struct {
	Body string
}

func main() {
	handler := rest.ResourceHandler{}
	handler.SetRoutes(
		&rest.Route{"GET", "/message", func(w rest.ResponseWriter, req *rest.Request) {
			w.WriteJson(&Message{
				Body: "Hello World!",
			})
		}},
	)
	http.Handle("/api/", http.StripPrefix("/api", &handler))

	http.Handle("/static/", http.StripPrefix("/static", http.FileServer(http.Dir("."))))

	http.ListenAndServe(":8080", nil)
}
