package gaehelloworld

import (
	"github.com/ant0ine/go-json-rest/rest"
	"log"
	"net/http"
)

func init() {
	handler := rest.ResourceHandler{}
	err := handler.SetRoutes(
		&rest.Route{"GET", "/message", func(w rest.ResponseWriter, req *rest.Request) {
			w.WriteJson(map[string]string{"Body": "Hello World!"})
		}},
	)
	if err != nil {
		log.Fatal(err)
	}
	http.Handle("/", &handler)
}
