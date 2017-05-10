package bank

import (
	"bitbucket.org/mundipagg/boletoapi/auth"
	"bitbucket.org/mundipagg/boletoapi/log"
	"bitbucket.org/mundipagg/boletoapi/models"
)

type bankCaixa struct {
	validate *models.Validator
	log      *log.Log
}

func newCaixa() bankCaixa {
	c := bankCaixa{
		validate: models.NewValidator(),
		log:      log.CreateLog(),
	}
	return c
}

//Log retorna a referencia do log
func (b bankCaixa) Log() *log.Log {
	return b.log
}
func (b bankCaixa) Login(user, password, body string) (auth.Token, error) {
	return auth.Token{Status: 200}, nil
}
func (b bankCaixa) RegisterBoleto(boleto *models.BoletoRequest) (models.BoletoResponse, error) {
	return models.BoletoResponse{}, nil
}
func (b bankCaixa) ProcessBoleto(boleto *models.BoletoRequest) (models.BoletoResponse, error) {
	return models.BoletoResponse{}, nil
}

func (b bankCaixa) ValidateBoleto(boleto *models.BoletoRequest) models.Errors {
	return nil
}

//GetBankNumber retorna o codigo do banco
func (b bankCaixa) GetBankNumber() models.BankNumber {
	return models.Caixa
}
