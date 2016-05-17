package main

import (
	"log"
	"net/http"
	"time"

	"github.com/StephanDollberg/go-json-rest-middleware-jwt"
	"github.com/ant0ine/go-json-rest/rest"
)

func handle_auth(w rest.ResponseWriter, r *rest.Request) {
	// w.WriteJson(map[string]string{"authed": r.Env["REMOTE_USER"].(string)})

	// Extract the JWT from the request environemnt
	jwt_claims := r.Env["JWT_PAYLOAD"].(map[string]interface{})

	// Build our custom response, write it
	response := map[string]interface{}{
		"authed":  r.Env["REMOTE_USER"],
		"isAdmin": jwt_claims["isAdmin"],
	}
	w.WriteJson(response)
}

func main() {
	jwt_middleware := &jwt.JWTMiddleware{
		Key:        []byte("secret key"),
		Realm:      "jwt auth",
		Timeout:    time.Hour,
		MaxRefresh: time.Hour * 24,

		// Authenticator is used to check user/pass combination
		Authenticator: func(userId string, password string) bool {
			isAdmin := (userId == "admin" && password == "admin")
			isUser := (userId == "user" && password == "pass")
			return (isAdmin || isUser)
		},

		// Authorizator checks that an authenticated user has access to
		// perform the given request
		Authorizator: func(userId string, request *rest.Request) bool {
			if request.RequestURI == "/api/admin_only" {
				if userId == "admin" {
					return true
				}
				return false
			}
			return true
		},

		// PayloadFunc allows us to append claims to the JWT that can later
		// be access
		PayloadFunc: func(userId string) map[string]interface{} {
			claims := make(map[string]interface{}, 0)
			claims["isAdmin"] = false
			if userId == "admin" {
				claims["isAdmin"] = true
			}
			return claims
		},
	}

	api := rest.NewApi()
	api.Use(rest.DefaultDevStack...)
	// we use the IfMiddleware to remove certain paths from needing authentication
	api.Use(&rest.IfMiddleware{
		Condition: func(request *rest.Request) bool {
			return request.URL.Path != "/login"
		},
		IfTrue: jwt_middleware,
	})
	api_router, _ := rest.MakeRouter(
		rest.Post("/login", jwt_middleware.LoginHandler),
		rest.Get("/auth_test", handle_auth),
		rest.Get("/admin_only", handle_auth), // Admin only route
		rest.Get("/refresh_token", jwt_middleware.RefreshHandler),
	)
	api.SetApp(api_router)

	http.Handle("/api/", http.StripPrefix("/api", api.MakeHandler()))

	log.Fatal(http.ListenAndServe(":8080", nil))
}
