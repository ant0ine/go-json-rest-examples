/* Demonstrate how to implement a custom AuthBasic middleware, used to protect all endpoints.

This is a very simple version supporting only one user.

The Curl Demo:

        curl -i http://127.0.0.1:8080/countries

*/
package main

import (
	"encoding/base64"
	"errors"
	"github.com/ant0ine/go-json-rest/rest"
	"net/http"
	"strings"
)

type MyAuthBasicMiddleware struct {
	Realm    string
	UserId   string
	Password string
}

func (mw *MyAuthBasicMiddleware) MiddlewareFunc(handler rest.HandlerFunc) rest.HandlerFunc {
	return func(writer rest.ResponseWriter, request *rest.Request) {

		authHeader := request.Header.Get("Authorization")
		if authHeader == "" {
			mw.unauthorized(writer)
			return
		}

		providedUserId, providedPassword, err := mw.decodeBasicAuthHeader(authHeader)

		if err != nil {
			rest.Error(writer, "Invalid authentication", http.StatusBadRequest)
			return
		}

		if !(providedUserId == mw.UserId && providedPassword == mw.Password) {
			mw.unauthorized(writer)
			return
		}

		handler(writer, request)
	}
}

func (mw *MyAuthBasicMiddleware) unauthorized(writer rest.ResponseWriter) {
	writer.Header().Set("WWW-Authenticate", "Basic realm="+mw.Realm)
	rest.Error(writer, "Not Authorized", http.StatusUnauthorized)
}

func (mw *MyAuthBasicMiddleware) decodeBasicAuthHeader(header string) (user string, password string, err error) {

	parts := strings.SplitN(header, " ", 2)
	if !(len(parts) == 2 && parts[0] == "Basic") {
		return "", "", errors.New("Invalid authentication")
	}

	decoded, err := base64.StdEncoding.DecodeString(parts[1])
	if err != nil {
		return "", "", errors.New("Invalid base64")
	}

	creds := strings.SplitN(string(decoded), ":", 2)
	if len(creds) != 2 {
		return "", "", errors.New("Invalid authentication")
	}

	return creds[0], creds[1], nil
}

func main() {

	handler := rest.ResourceHandler{
		PreRoutingMiddlewares: []rest.Middleware{
			&MyAuthBasicMiddleware{
				Realm:    "Administration",
				UserId:   "admin",
				Password: "admin",
			},
		},
	}
	handler.SetRoutes(
		rest.Route{"GET", "/countries", GetAllCountries},
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
