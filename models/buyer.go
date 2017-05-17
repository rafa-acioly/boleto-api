package models

// Buyer informações de entrada do comprador
type Buyer struct {
	Name     string   `json:"name,omitempty"`
	Document Document `json:"document,omitempty"`
	Address  Address  `json:"address,omitempty"`
}
