package gaehelloworld

import (
	"github.com/ant0ine/go-json-rest/rest"
	"log"
	"net/http"
)

type Message struct {
	Body string
}

func init() {
	handler := rest.ResourceHandler{}
	err := handler.SetRoutes(
		&rest.Route{"GET", "/message", func(w rest.ResponseWriter, req *rest.Request) {
			w.WriteJson(&Message{
				Body: "Hello World!",
			})
		}},
	)
	if err != nil {
		log.Fatal(err)
	}
	http.Handle("/", &handler)
}
