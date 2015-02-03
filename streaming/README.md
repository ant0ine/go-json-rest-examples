# Streaming

Demonstrate a streaming REST API, where the data is "flushed" to the client ASAP.

The stream format is a Line Delimited JSON.

curl demo:
```
curl -i http://127.0.0.1:8080/stream
```

Output:
```
HTTP/1.1 200 OK
Content-Type: application/json
Date: Sun, 16 Feb 2014 00:39:19 GMT
Transfer-Encoding: chunked

{"Name":"thing #1"}
{"Name":"thing #2"}
{"Name":"thing #3"}
```
