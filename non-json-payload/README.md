# Non JSON payload

Exceptional use of non JSON payloads.

The ResponseWriter implementation provided by go-json-rest is designed
to build JSON responses. In order to serve different kind of content,
it is recommended to either:
a) use another server and configure CORS
   (see the cors/ example)
b) combine the api.MakeHandler() with another http.Handler
   (see api-and-static/ example)

That been said, exceptionally, it can be convenient to return a
different content type on a JSON endpoint. In this case, setting the
Content-Type and using the type assertion to access the Write method
is enough. As shown in this example.

curl demo:
```
curl -i http://127.0.0.1:8080/message.txt
```
