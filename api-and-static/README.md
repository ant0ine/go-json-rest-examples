# API and static files

Combine Go-Json-Rest with other handlers.

`api.MakeHandler()` is a valid `http.Handler`, and can be combined with other handlers.
In this example the api handler is used under the `/api/` prefix, while a FileServer is instantiated under the `/static/` prefix.

curl demo:
```
curl -i http://127.0.0.1:8080/api/message
curl -i http://127.0.0.1:8080/static/main.go
```
