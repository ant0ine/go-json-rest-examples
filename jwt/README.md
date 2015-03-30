# JWT

Demonstrates how to use the [Json Web Token Auth Middleware](https://github.com/StephanDollberg/go-json-rest-middleware-jwt) to authenticate via a JWT token.

curl demo:
``` sh
curl -d '{"username": "admin", "password": "admin"}' -H "Content-Type:application/json" http://localhost:8080/api/login
curl -H "Authorization:Bearer TOKEN_RETURNED_FROM_ABOVE" http://localhost:8080/api/auth_test
curl -H "Authorization:Bearer TOKEN_RETURNED_FROM_ABOVE" http://localhost:8080/api/refresh_token
```
