package bank

import (
	"fmt"

	"bitbucket.org/mundipagg/boletoapi/auth"
	"bitbucket.org/mundipagg/boletoapi/log"
	"bitbucket.org/mundipagg/boletoapi/models"
)

//Bank é a interface que vai oferecer os serviços em comum entre os bancos
type Bank interface {
	Login(string, string, string) (auth.Token, error)
	RegisterBoleto(models.BoletoRequest) (string, error)
	ValidateBoleto(models.BoletoRequest) []string
	GetBankNumber() models.BankNumber
	Log() *log.Log
}

//Get retorna estrategia de acordo com o banco ou erro caso o banco não exista
func Get(number models.BankNumber) (Bank, error) {
	switch number {
	case models.BancoDoBrasil:
		return bankBB{
			log: log.CreateLog(),
		}, nil
	case models.Caixa:
		return bankCaixa{
			log: log.CreateLog(),
		}, nil
	default:
		return nil, fmt.Errorf("Banco %d não existe", number)
	}
}
