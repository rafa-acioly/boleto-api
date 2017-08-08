package models

// Recipient informações de entrada do comprador
type Recipient struct {
	Name     string   `json:"name,omitempty"`
	Document Document `json:"document,omitempty"`
	Address  Address  `json:"address,omitempty"`
}
