package socketio

import "github.com/gorilla/websocket"

type (
	WebSocketConnection struct {
		*websocket.Conn
	}

	WsJsonResponse struct {
		Action      string `json:"action"`
		Message     string `json:"message"`
		MessageType string `json:"message_type"`
	}

	WsPayload struct {
		Action   string              `json:"action"`
		Username string              `json:"username"`
		Message  string              `json:"message"`
		Conn     WebSocketConnection `json:"-"`
	}
)
