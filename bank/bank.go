package bank

import (
	"fmt"

	"bitbucket.org/mundipagg/boletoapi/auth"
	"bitbucket.org/mundipagg/boletoapi/models"
)

//Bank é a interface que vai oferecer os serviços em comum entre os bancos
type Bank interface {
	Login() auth.Token
	RegisterBoleto(models.BoletoRequest) models.BoletoResponse
	GetBankNumber() models.BankNumber
}

var bankRouter map[models.BankNumber]Bank

//InstallBanks instala os bancos configurados no "roteador" de bancos
func InstallBanks() {
	bankRouter = make(map[models.BankNumber]Bank)
	bankRouter[models.BancoDoBrasil] = bankBB{}
}

//Get retornaa estrategia de acordo com o banco ou erro caso o banco não exista
func Get(number models.BankNumber) (Bank, error) {
	bank, ok := bankRouter[number]
	if !ok {
		return nil, fmt.Errorf("Banco %d não existe", number)
	}
	return bank, nil
}
