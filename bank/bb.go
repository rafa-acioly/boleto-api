package bank

import (
	"bitbucket.org/mundipagg/boletoapi/auth"
	"bitbucket.org/mundipagg/boletoapi/models"
)

type bankBB struct {
}

func (b bankBB) Login() auth.Token {
	return auth.Token{}
}
func (b bankBB) RegisterBoleto(boleto models.BoletoRequest) models.BoletoResponse {
	//fmt.Println("Registrando boleto no Banco do Brasil")

	return models.BoletoResponse{}
}
func (b bankBB) GetBankNumber() models.BankNumber {
	return models.BancoDoBrasil
}
