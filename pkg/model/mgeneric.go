package model

// Message model to report a message
type Message struct {
	Message string `json:"message"`
	ID      string `json:"identifier"`
}
