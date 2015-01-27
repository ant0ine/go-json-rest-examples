package main

import (
	"github.com/ant0ine/go-json-rest/rest"
	"log"
	"net/http"
)

func main() {
	router, err := rest.MakeRouter(
		&rest.Route{"GET", "/countries", GetAllCountries},
	)
	if err != nil {
		log.Fatal(err)
	}

	api := rest.NewApi(router)
	api.Use(rest.DefaultDevStack...)
	api.Use(&rest.AuthBasicMiddleware{
		Realm: "test zone",
		Authenticator: func(userId string, password string) bool {
			if userId == "admin" && password == "admin" {
				return true
			}
			return false
		},
	})
	log.Fatal(http.ListenAndServe(":8080", api.MakeHandler()))
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
