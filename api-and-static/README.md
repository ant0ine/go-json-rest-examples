# API and static files

Combine Go-Json-Rest with other handlers.

`rest.ResourceHandler` is a valid `http.Handler`, and can be combined with other handlers.
In this example the ResourceHandler is used under the `/api/` prefix, while a FileServer is instantiated under the `/static/` prefix.

The curl demo:

        curl -i http://127.0.0.1:8080/api/message
        curl -i http://127.0.0.1:8080/static/main.go

