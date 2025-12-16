package models

// Response é um modelo básico para nossas respostas JSON
type Response struct {
	Message string `json:"message,omitempty"`
	Value   any    `json:"value,omitempty"`
}
