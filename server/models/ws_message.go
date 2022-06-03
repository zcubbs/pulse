package models

type WSMessage struct {
	MessageType string `json:"message_type"`
	Message     string `json:"message"`
}
