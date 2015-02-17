package main

import (
	"github.com/StephanDollberg/go-json-rest-middleware-jwt"
	"github.com/ant0ine/go-json-rest/rest"
	"log"
	"net/http"
	"time"
)

func handle_auth(w rest.ResponseWriter, r *rest.Request) {
	w.WriteJson(map[string]string{"authed": r.Env["REMOTE_USER"].(string)})
}

func main() {
	jwt_middleware := jwt.JWTMiddleware{
		Key:        []byte("secret key"),
		Realm:      "jwt auth",
		Timeout:    time.Hour,
		MaxRefresh: time.Hour * 24,
		Authenticator: func(userId string, password string) bool {
			if userId == "admin" && password == "admin" {
				return true
			}
			return false
		}}

	login_api := rest.NewApi()
	login_api.Use(rest.DefaultDevStack...)
	login_router, _ := rest.MakeRouter(
		&rest.Route{"POST", "/login", jwt_middleware.LoginHandler},
	)
	login_api.SetApp(login_router)

	main_api := rest.NewApi()
	main_api.Use(&jwt_middleware)
	main_api.Use(rest.DefaultDevStack...)
	main_api_router, _ := rest.MakeRouter(
		&rest.Route{"GET", "/auth_test", handle_auth},
		&rest.Route{"GET", "/refresh_token", jwt_middleware.RefreshHandler})
	main_api.SetApp(main_api_router)

	http.Handle("/", login_api.MakeHandler())
	http.Handle("/api/", http.StripPrefix("/api", main_api.MakeHandler()))

	log.Fatal(http.ListenAndServe(":8080", nil))
}
