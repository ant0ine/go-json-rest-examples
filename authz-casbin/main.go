package main

import (
	"github.com/ant0ine/go-json-rest/rest"
	"github.com/casbin/casbin"

	"log"
	"net/http"
)

func main() {
	e := casbin.NewEnforcer("authz-casbin/authz_model.conf", "authz-casbin/authz_policy.csv")

	api := rest.NewApi()
	api.Use(rest.DefaultDevStack...)
	api.Use(&rest.AuthBasicMiddleware{
		Realm: "test zone",
		Authenticator: func(userId string, password string) bool {
			if userId == "alice" && password == "123" {
				return true
			} else if userId == "bob" && password == "123" {
				return true
			} else if userId == "cathy" && password == "123" {
				return true
			}
			return false
		},
		Authorizator: func(userId string, request *rest.Request) bool {
			method := request.Method
			path := request.URL.Path

			return e.Enforce(userId, path, method)
		},
	})
	api.SetApp(rest.AppSimple(func(w rest.ResponseWriter, r *rest.Request) {
		w.WriteJson(map[string]string{"Body": "Hello World!"})
	}))
	log.Fatal(http.ListenAndServe(":8080", api.MakeHandler()))
}
