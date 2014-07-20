package main

import (
	"fmt"
	"github.com/ant0ine/go-json-rest/rest"
	"log"
	"net/http"
	"sync"
)

func main() {

	users := Users{
		Store: map[string]*User{},
	}

	handler := rest.ResourceHandler{
		EnableRelaxedContentType: true,
	}
	err := handler.SetRoutes(
		rest.RouteObjectMethod("GET", "/users", &users, "GetAllUsers"),
		rest.RouteObjectMethod("POST", "/users", &users, "PostUser"),
		rest.RouteObjectMethod("GET", "/users/:id", &users, "GetUser"),
		rest.RouteObjectMethod("PUT", "/users/:id", &users, "PutUser"),
		rest.RouteObjectMethod("DELETE", "/users/:id", &users, "DeleteUser"),
	)
	if err != nil {
		log.Fatal(err)
	}
	log.Fatal(http.ListenAndServe(":8080", &handler))
}

type User struct {
	Id   string
	Name string
}

type Users struct {
	sync.RWMutex
	Store map[string]*User
}

func (u *Users) GetAllUsers(w rest.ResponseWriter, r *rest.Request) {
	u.RLock()
	users := make([]User, len(u.Store))
	i := 0
	for _, user := range u.Store {
		users[i] = *user
		i++
	}
	u.RUnlock()
	w.WriteJson(&users)
}

func (u *Users) GetUser(w rest.ResponseWriter, r *rest.Request) {
	id := r.PathParam("id")
	u.RLock()
	var user *User
	if u.Store[id] != nil {
		user = &User{}
		*user = *u.Store[id]
	}
	u.RUnlock()
	if user == nil {
		rest.NotFound(w, r)
		return
	}
	w.WriteJson(user)
}

func (u *Users) PostUser(w rest.ResponseWriter, r *rest.Request) {
	user := User{}
	err := r.DecodeJsonPayload(&user)
	if err != nil {
		rest.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	u.Lock()
	id := fmt.Sprintf("%d", len(u.Store)) // stupid
	user.Id = id
	u.Store[id] = &user
	u.Unlock()
	w.WriteJson(&user)
}

func (u *Users) PutUser(w rest.ResponseWriter, r *rest.Request) {
	id := r.PathParam("id")
	u.Lock()
	if u.Store[id] == nil {
		rest.NotFound(w, r)
		u.Unlock()
		return
	}
	user := User{}
	err := r.DecodeJsonPayload(&user)
	if err != nil {
		rest.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	user.Id = id
	u.Store[id] = &user
	u.Unlock()
	w.WriteJson(&user)
}

func (u *Users) DeleteUser(w rest.ResponseWriter, r *rest.Request) {
	id := r.PathParam("id")
	u.Lock()
	delete(u.Store, id)
	u.Unlock()
	w.WriteHeader(http.StatusOK)
}
