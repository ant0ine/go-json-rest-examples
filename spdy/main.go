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
	handler := rest.ResourceHandler{}
	err := handler.SetRoutes(
		&rest.Route{"GET", "/users/:id", GetUser},
	)
	if err != nil {
		log.Fatal(err)
	}
	log.Fatal(spdy.ListenAndServeTCP(":8080", &handler))
}
