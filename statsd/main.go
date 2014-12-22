package main

import (
	"github.com/ant0ine/go-json-rest/rest"
	"github.com/peterbourgon/g2s"
	"log"
	"net/http"
	"strconv"
	"time"
)

type StatsdMiddleware struct {
	IpPort string
	Prefix string
}

func (mw *StatsdMiddleware) MiddlewareFunc(handler rest.HandlerFunc) rest.HandlerFunc {

	statsd, err := g2s.Dial("udp", mw.IpPort)
	if err != nil {
		panic(err)
	}

	keyBase := ""
	if mw.Prefix != "" {
		keyBase += mw.Prefix + "."
	}
	keyBase += "response."

	return func(writer rest.ResponseWriter, request *rest.Request) {

		handler(writer, request)

		statusCode := request.Env["STATUS_CODE"].(int)
		statsd.Counter(1.0, keyBase+"status_code."+strconv.Itoa(statusCode), 1)

		elapsedTime := request.Env["ELAPSED_TIME"].(*time.Duration)
		statsd.Timing(1.0, keyBase+"elapsed_time", *elapsedTime)
	}
}

func main() {
	handler := rest.ResourceHandler{
		OuterMiddlewares: []rest.Middleware{
			&StatsdMiddleware{
				IpPort: "localhost:8125",
			},
		},
	}
	err := handler.SetRoutes(
		&rest.Route{"GET", "/message", func(w rest.ResponseWriter, req *rest.Request) {

			// take more than 1ms so statsd can report it
			time.Sleep(100 * time.Millisecond)

			w.WriteJson(map[string]string{"Body": "Hello World!"})
		}},
	)
	if err != nil {
		log.Fatal(err)
	}
	log.Fatal(http.ListenAndServe(":8080", &handler))
}
