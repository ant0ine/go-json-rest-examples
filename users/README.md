# Users

Demonstrate how to use Method Values.

Method Values have been [introduced in Go 1.1](https://golang.org/doc/go1.1#method_values).

This shows how to map a Route to a method of an instantiated object (i.e: receiver of the method)

curl demo:
```
curl -i -H 'Content-Type: application/json' \
    -d '{"Name":"Antoine"}' http://127.0.0.1:8080/users
curl -i http://127.0.0.1:8080/users/0
curl -i -X PUT -H 'Content-Type: application/json' \
    -d '{"Name":"Antoine Imbert"}' http://127.0.0.1:8080/users/0
curl -i -X DELETE http://127.0.0.1:8080/users/0
curl -i http://127.0.0.1:8080/users
```
