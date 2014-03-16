/* Demonstrate a possible global Basic Auth implementation

The Curl Demo:

        curl -i http://127.0.0.1:8080/countries

*/
package main

import (
	"encoding/base64"
	"errors"
	"github.com/ant0ine/go-json-rest"
	"net/http"
	"strings"
)

func main() {

	handler := rest.ResourceHandler{
		PreRoutingMiddleware: func(handler rest.HandlerFunc) rest.HandlerFunc {

			realm := "Administration"
			userId := "admin"
			password := "admin"

			return func(writer rest.ResponseWriter, request *rest.Request) {

				authHeader := request.Header.Get("Authorization")
				if authHeader == "" {
					Unauthorized(writer, realm)
					return
				}

				providedUserId, providedPassword, err := DecodeBasicAuthHeader(authHeader)

				if err != nil {
					rest.Error(writer, "Invalid authentication", http.StatusBadRequest)
					return
				}

				if !(providedUserId == userId && providedPassword == password) {
					Unauthorized(writer, realm)
					return
				}

				handler(writer, request)
			}
		},
	}
	handler.SetRoutes(
		rest.Route{"GET", "/countries", GetAllCountries},
	)
	http.ListenAndServe(":8080", &handler)
}

func Unauthorized(writer rest.ResponseWriter, realm string) {
	writer.Header().Set("WWW-Authenticate", "Basic realm="+realm)
	rest.Error(writer, "Not Authorized", http.StatusUnauthorized)
}

func DecodeBasicAuthHeader(header string) (user string, password string, err error) {

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
