package main

import (
	"github.com/ant0ine/go-json-rest/rest"
	"github.com/jadengore/go-json-rest-middleware-force-ssl"
	"log"
	"net/http"
)

func main() {
	api := rest.NewApi()
	api.Use(&forceSSL.Middleware{
		TrustXFPHeader:     true,
		Enable301Redirects: false,
	})
	api.SetApp(rest.AppSimple(func(w rest.ResponseWriter, r *rest.Request) {
		w.WriteJson(map[string]string{"body": "Hello World!"})
	}))

	// For the purposes of this demo, only HTTP connections accepted.
	// For true HTTPS, use ListenAndServeTLS.
	// https://golang.org/pkg/net/http/#ListenAndServeTLS
	log.Fatal(http.ListenAndServe(":8080", api.MakeHandler()))
}
