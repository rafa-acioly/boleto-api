package models

// BoletoRequest entidade de entrada para o boleto
type BoletoRequest struct {
	Authentication Authentication
	Agreement      Agreement
	Title          Title
	Buyer          Buyer
	BankNumber     BankNumber
}

// BoletoResponse entidade de saída para o boleto
type BoletoResponse struct {
	StatusCode       int `json:"-"`
	Error            string
	ErrorDescription string
}

// BankNumber número de identificação do banco
type BankNumber int

// IsBankNumberValid verifica se o banco enviado existe
func (b BankNumber) IsBankNumberValid() bool {
	switch b {
	case BancoDoBrasil, Itau, Santander, Caixa, Bradesco:
		return true
	default:
		return false
	}
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

	// Caixa constante do Caixa
	Caixa = 104
)
