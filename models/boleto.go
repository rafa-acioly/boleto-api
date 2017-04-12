package models

// BoletoRequest entidade de entrada para o boleto
type BoletoRequest struct {
	Authentication Authentication
	Agreement      Agreement
	Title          Title
	Buyer          Buyer
	BankNumber     int
}

// BoletoResponse entidade de saída para o boleto
type BoletoResponse struct {
}

const (
	// BancoDoBrasil constante do Banco do Brasil
	BancoDoBrasil = 1

	// Santander constante do Santander
	Santander = 33

	// Itau constante do Itau
	Itau = 341

	// Bradesco constante do Bradesco
	Bradesco = 237
)
