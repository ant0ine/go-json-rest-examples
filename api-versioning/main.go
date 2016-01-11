package main

import (
	"github.com/ant0ine/go-json-rest/rest"
	"github.com/coreos/go-semver/semver"
	"log"
	"net/http"
)

type SemVerMiddleware struct {
	MinVersion string
	MaxVersion string
}

func (mw *SemVerMiddleware) MiddlewareFunc(handler rest.HandlerFunc) rest.HandlerFunc {

	minVersion, err := semver.NewVersion(mw.MinVersion)
	if err != nil {
		panic(err)
	}

	maxVersion, err := semver.NewVersion(mw.MaxVersion)
	if err != nil {
		panic(err)
	}

	return func(writer rest.ResponseWriter, request *rest.Request) {

		version, err := semver.NewVersion(request.PathParam("version"))
		if err != nil {
			rest.Error(
				writer,
				"Invalid version: "+err.Error(),
				http.StatusBadRequest,
			)
			return
		}

		if version.LessThan(*minVersion) {
			rest.Error(
				writer,
				"Min supported version is "+minVersion.String(),
				http.StatusBadRequest,
			)
			return
		}

		if maxVersion.LessThan(*version) {
			rest.Error(
				writer,
				"Max supported version is "+maxVersion.String(),
				http.StatusBadRequest,
			)
			return
		}

		request.Env["VERSION"] = version
		handler(writer, request)
	}
}

func main() {

	svmw := SemVerMiddleware{
		MinVersion: "1.0.0",
		MaxVersion: "3.0.0",
	}
	api := rest.NewApi()
	api.Use(rest.DefaultDevStack...)
	router, err := rest.MakeRouter(
		rest.Get("/#version/message", svmw.MiddlewareFunc(
			func(w rest.ResponseWriter, req *rest.Request) {
				version := req.Env["VERSION"].(*semver.Version)
				if version.Major == 2 {
					// https://en.wikipedia.org/wiki/Second-system_effect
					w.WriteJson(map[string]string{
						"Body": "Hello broken World!",
					})
				} else {
					w.WriteJson(map[string]string{
						"Body": "Hello World!",
					})
				}
			},
		)),
	)
	if err != nil {
		log.Fatal(err)
	}
	api.SetApp(router)
	http.Handle("/api/", http.StripPrefix("/api", api.MakeHandler()))
	log.Fatal(http.ListenAndServe(":8080", nil))
}
