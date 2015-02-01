# API Versioning

First, API versioning is not easy and you may want to favor a mechanism that uses only backward compatible changes and deprecation cycles.

That been said, here is an example of API versioning using [Semver](http://semver.org/)

It defines a middleware that parses the version, checks a min and a max, and makes it available in the `request.Env`.

curl demo:
``` sh
curl -i http://127.0.0.1:8080/api/1.0.0/message
curl -i http://127.0.0.1:8080/api/2.0.0/message
curl -i http://127.0.0.1:8080/api/2.0.1/message
curl -i http://127.0.0.1:8080/api/0.0.1/message
curl -i http://127.0.0.1:8080/api/4.0.1/message

```
