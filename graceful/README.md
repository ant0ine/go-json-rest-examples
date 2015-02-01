# Graceful Shutdown

This example uses [github.com/stretchr/graceful](https://github.com/stretchr/graceful) to try to be nice with the clients waiting for responses during a server shutdown (or restart).
The HTTP response takes 10 seconds to be completed, printing a message on the wire every second.
10 seconds is also the timeout set for the graceful shutdown.
You can play with these numbers to show that the server waits for the responses to complete.

curl demo:
``` sh
curl -i http://127.0.0.1:8080/message
```
