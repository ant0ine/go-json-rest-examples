/* Demonstrate how to setup CorsMiddleware around all the API endpoints.

The curl demo:

        curl -i http://127.0.0.1:8080/countries

*/
package main

import (
	"github.com/ant0ine/go-json-rest/rest"
	"net/http"
)

func main() {

	handler := rest.ResourceHandler{
		PreRoutingMiddlewares: []rest.Middleware{
			&rest.CorsMiddleware{
				RejectNonCorsRequests: false,
				OriginValidator: func(origin string, request *rest.Request) bool {
					return origin == "http://my.other.host"
				},
				AllowedMethods: []string{"GET", "POST", "PUT"},
				AllowedHeaders: []string{
					"Accept", "Content-Type", "X-Custom-Header", "Origin"},
				AccessControlAllowCredentials: true,
				AccessControlMaxAge:           3600,
			},
		},
	}
	handler.SetRoutes(
		&rest.Route{"GET", "/countries", GetAllCountries},
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
