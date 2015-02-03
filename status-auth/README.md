# Status Auth

Demonstrate how to setup a /.status endpoint protected with basic authentication.

This is a good use case of middleware applied to only one API endpoint.

curl demo:
```
curl -i http://127.0.0.1:8080/countries
curl -i http://127.0.0.1:8080/.status
curl -i -u admin:admin http://127.0.0.1:8080/.status
...
```
