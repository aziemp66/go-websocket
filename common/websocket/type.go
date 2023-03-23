package socketio

import "github.com/gorilla/websocket"

type (
	WebSocketConnection struct {
		*websocket.Conn
	}

	// WsJsonResponse defines the response sent back from websocket
	WsJsonResponse struct {
		Action      string `json:"action"`
		Message     string `json:"message"`
		MessageType string `json:"message_type"`
	}

	// WsPayload defines the websocket request from the client
	WsPayload struct {
		Action   string              `json:"action"`
		Username string              `json:"username"`
		Message  string              `json:"message"`
		Conn     WebSocketConnection `json:"-"`
	}
)
