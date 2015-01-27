package main

import (
	"github.com/ant0ine/go-json-rest/rest"
	"github.com/shykes/spdy-go"
	"log"
)

type User struct {
	Id   string
	Name string
}

func GetUser(w rest.ResponseWriter, req *rest.Request) {
	user := User{
		Id:   req.PathParam("id"),
		Name: "Antoine",
	}
	w.WriteJson(&user)
}

func main() {
	router, err := rest.MakeRouter(
		&rest.Route{"GET", "/users/:id", GetUser},
	)
	if err != nil {
		log.Fatal(err)
	}

	api := rest.NewApi(router)
	api.Use(rest.DefaultDevStack...)
	log.Fatal(spdy.ListenAndServeTCP(":8080", api.MakeHandler()))
}
