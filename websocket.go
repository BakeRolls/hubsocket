package hubsocket

import (
	"golang.org/x/net/websocket"
)

var clients []*websocket.Conn
var listeners map[string][]func(*websocket.Conn, string)

func init() {
	clients = []*websocket.Conn{}
	listeners = make(map[string][]func(*websocket.Conn, string))
}

func add(ws *websocket.Conn) {
	clients = append(clients, ws)
}

func delete(ws *websocket.Conn) {
	for i, c := range clients {
		if c == ws {
			clients[i], clients[len(clients)-1] = clients[len(clients)-1], clients[i]
			clients = clients[:len(clients)-1]
			break
		}
	}
}

func call(event string, ws *websocket.Conn, body string) {
	for _, cb := range listeners[event] {
		cb(ws, body)
	}
}

// Send a message to a specific client
func Send(ws *websocket.Conn, m Message) {
	websocket.JSON.Send(ws, m)
}

// Broadcast a message to all connected clients
func Broadcast(m Message) {
	for _, ws := range clients {
		go Send(ws, m)
	}
}

// Clients returns the number of active connections
func Clients() int {
	return len(clients)
}

// Handle a callback
func Handle(event string, cb func(*websocket.Conn, string)) {
	listeners[event] = append(listeners[event], cb)
}

func handler(ws *websocket.Conn) {
	add(ws)
	call("connect", ws, "")

	for {
		m := Message{}
		if err := websocket.JSON.Receive(ws, &m); err != nil {
			delete(ws)
			call("disconnect", ws, "")
			break
		}
		call(m.Event, ws, m.Body)
	}
}

// Handler for websockets
func Handler() websocket.Handler {
	return websocket.Handler(handler)
}
