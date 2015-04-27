# Websocket

Demonstrate how to run websocket in go-json-rest

go client demo:
```go
origin := "http://localhost:8080/"
url := "ws://localhost:8080/ws"
ws, err := websocket.Dial(url, "", origin)
if err != nil {
	log.Fatal(err)
}
if _, err := ws.Write([]byte("hello, world\n")); err != nil {
	log.Fatal(err)
}
var msg = make([]byte, 512)
var n int
if n, err = ws.Read(msg); err != nil {
	log.Fatal(err)
}
log.Printf("Received: %s.", msg[:n])
```
