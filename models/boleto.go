package models

// BoletoRequest entidade de entrada para o boleto
type BoletoRequest struct {
	Authentication Authentication
	Agreement      Agreement
	Title          Title
	Buyer          Buyer
}

// BoletoResponse entidade de saída para o boleto
type BoletoResponse struct {
}
