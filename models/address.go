package models

// Address informações de entrada do endereço
type Address struct {
	Street     string
	Number     string
	Complement string
	ZipCode    string
	City       string
	District   string
	StateCode  string
}
