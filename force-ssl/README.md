# ForceSSL

Demonstrate how to use the [ForceSSL Middleware](https://github.com/jadengore/go-json-rest-middleware-force-ssl) to force HTTPS on requests to a `go-json-rest` API.

For the purposes of this demo, we are using HTTP for all requests and checking the `X-Forwarded-Proto` header to see if it is set to HTTPS (many routers set this to show what type of connection the client is using, such as Heroku). To do a true HTTPS test, make sure and use [`http.ListenAndServeTLS`](https://golang.org/pkg/net/http/#ListenAndServeTLS) with a valid certificate and key file.

Additional documentation for the ForceSSL middleware can be found [here](https://github.com/jadengore/go-json-rest-middleware-force-ssl).

curl demo:
``` sh
curl -i 127.0.0.1:8080/
curl -H "X-Forwarded-Proto:https" -i 127.0.0.1:8080/
```
