package gaehelloworld

import (
	"github.com/ant0ine/go-json-rest/rest"
	"net/http"
)

type Message struct {
	Body string
}

func init() {
	handler := rest.ResourceHandler{}
	handler.SetRoutes(
		&rest.Route{"GET", "/message", func(w rest.ResponseWriter, req *rest.Request) {
			w.WriteJson(&Message{
				Body: "Hello World!",
			})
		}},
	)
	http.Handle("/", &handler)
}
