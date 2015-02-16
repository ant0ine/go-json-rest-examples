# Statsd

Demonstrate how to use the [Statsd Middleware](https://github.com/ant0ine/go-json-rest-middleware-statsd) to collect statistics about the requests/reponses.
This middleware is based on the [g2s](https://github.com/peterbourgon/g2s) statsd client.

curl demo:
``` sh
# start statsd server
# monitor network
ngrep -d any port 8125

curl -i http://127.0.0.1:8080/message
curl -i http://127.0.0.1:8080/doesnotexist

```
