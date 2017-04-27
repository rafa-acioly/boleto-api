package models

import (
	"html/template"

	"github.com/google/uuid"

	"fmt"

	"encoding/json"

	"bitbucket.org/mundipagg/boletoapi/config"
)

// BoletoRequest entidade de entrada para o boleto
type BoletoRequest struct {
	Authentication Authentication
	Agreement      Agreement
	Title          Title
	Recipient      Recipient
	Buyer          Buyer
	BankNumber     BankNumber
}

// BoletoResponse entidade de saída para o boleto
type BoletoResponse struct {
	StatusCode    int    `json:"-"`
	Errors        Errors `json:",omitempty"`
	URL           string `json:"Url,omitempty"`
	DigitableLine string `json:",omitempty"`
	BarCodeNumber string `json:",omitempty"`
}

// BoletoView contem as informações que serão preenchidas no boleto
type BoletoView struct {
	BankLogo      template.HTML
	Boleto        BoletoRequest
	BankID        BankNumber
	BankNumber    string
	DigitableLine string
	Barcode       string
	Barcode64     string
}

// NewBoletoView cria um novo objeto view de boleto a partir de um boleto request, codigo de barras e linha digitavel
func NewBoletoView(boleto BoletoRequest, barcode string, digitableLine string) BoletoView {
	view := BoletoView{
		BankID:        boleto.BankNumber,
		Boleto:        boleto,
		Barcode:       barcode,
		DigitableLine: digitableLine,
		BankNumber:    boleto.BankNumber.GetBoletoBankNumberAndDigit(),
	}
	return view
}

//EncodeURL tranforma o boleto view na forma que será escrito na url
func (b BoletoView) EncodeURL() (string, string) {
	id, _ := uuid.NewUUID()
	url := fmt.Sprintf("%s?fmt=html&id=%s", config.Get().AppURL, id)
	return url, id.String()
}

//ToJSON tranforma o boleto view em json
func (b BoletoView) ToJSON() string {
	json, _ := json.Marshal(b)
	return string(json)
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

func (b BankNumber) GetBoletoBankNumberAndDigit() string {
	switch b {
	case BancoDoBrasil:
		return "001-9"
	case Caixa:
		return "104-0"
	case Santander:
		return "033-7"
	case Itau:
		return "341-7"
	case Bradesco:
		return "237-2"
	default:
		return ""
	}
}

// BankName retorna o nome do banco
func (b BankNumber) BankName() string {
	switch b {
	case BancoDoBrasil:
		return "BancoDoBrasil"
	case Itau:
		return "Itau"
	case Santander:
		return "Santander"
	case Caixa:
		return "Caixa"
	case Bradesco:
		return "Bradesco"
	default:
		return ""
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
