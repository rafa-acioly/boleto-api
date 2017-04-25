package models

// Recipient informações de entrada do comprador
type Recipient struct {
	Name     string
	Document Document
	Address  Address
}
