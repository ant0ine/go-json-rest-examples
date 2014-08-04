# Statsd

Demonstrate how to use OuterMiddlewares to do additional logging and reporting.

Here `request.Env["STATUS_CODE"]` and `request.Env["ELAPSED_TIME"]` that are available to outer middlewares are used with the [g2s](https://github.com/peterbourgon/g2s) statsd client to send these metrics to statsd.

The curl demo:
``` sh
# start statsd server
# monitor network
ngrep -d any port 8125

curl -i http://127.0.0.1:8080/message
curl -i http://127.0.0.1:8080/doesnotexist

```
