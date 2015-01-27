package gaehelloworld

import (
	"github.com/ant0ine/go-json-rest/rest"
	"log"
	"net/http"
)

func init() {
	router, err := rest.MakeRouter(
		&rest.Route{"GET", "/message", func(w rest.ResponseWriter, req *rest.Request) {
			w.WriteJson(map[string]string{"Body": "Hello World!"})
		}},
	)
	if err != nil {
		log.Fatal(err)
	}

	api := rest.NewApi(router)
	api.Use(rest.DefaultDevStack...)
	http.Handle("/", api.MakeHandler())
}
