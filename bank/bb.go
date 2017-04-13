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
	return models.BoletoResponse{}
}
