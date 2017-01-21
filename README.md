# hubsocket

```go
package main

import (
	"fmt"
	"git.192k.pw/bake/hubsocket"
	"golang.org/x/net/websocket"
	"net/http"
)

func main() {
	hubsocket.Handle("connect", func(ws *websocket.Conn, body string) {
		fmt.Printf("%d clients\n", hubsocket.Clients())
	})
	http.Handle("/", hubsocket.Handler())
	http.ListenAndServe(":8080", nil)
}

```
