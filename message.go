package hubsocket

import (
	"encoding/json"
)

// Message with metadata
type Message struct {
	Event string `json:"event"`
	Body  string `json:"body"`
}

func (m *Message) String() string {
	message, _ := json.Marshal(m)
	return string(message)
}
