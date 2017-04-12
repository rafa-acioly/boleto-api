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
}

// BankNumber número de identificação do banco
type BankNumber int16

// IsBankNumberValid verifica se o banco enviado existe
func (b BankNumber) IsBankNumberValid() bool {
	switch int16(b) {
	case Santander:
	case Itau:
	case Bradesco:
	case Caixa:
	case BancoDoBrasil:
		return true
	default:
		return false
	}
	return false
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
