/* Demonstrate how to setup a /.status endpoint protected with basic authentication.

This is a good use case of middleware applied to only one API endpoint.

The Curl Demo:

        curl -i http://127.0.0.1:8080/countries
        curl -i http://127.0.0.1:8080/.status
        curl -i -u admin:admin http://127.0.0.1:8080/.status
        ...

*/
package main

import (
	"github.com/ant0ine/go-json-rest/rest"
	"net/http"
)

func main() {
	handler := rest.ResourceHandler{
		EnableStatusService: true,
	}
	auth := &rest.AuthBasicMiddleware{
		Realm: "test zone",
		Authenticator: func(userId string, password string) bool {
			if userId == "admin" && password == "admin" {
				return true
			}
			return false
		},
	}
	handler.SetRoutes(
		rest.Route{"GET", "/countries", GetAllCountries},
		rest.Route{"GET", "/.status",
			auth.MiddlewareFunc(
				func(w rest.ResponseWriter, r *rest.Request) {
					w.WriteJson(handler.GetStatus())
				},
			),
		},
	)
	http.ListenAndServe(":8080", &handler)
}

type Country struct {
	Code string
	Name string
}

func GetAllCountries(w rest.ResponseWriter, r *rest.Request) {
	w.WriteJson(
		[]Country{
			Country{
				Code: "FR",
				Name: "France",
			},
			Country{
				Code: "US",
				Name: "United States",
			},
		},
	)
}
