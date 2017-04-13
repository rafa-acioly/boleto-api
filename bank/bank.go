package bank

import (
	"bitbucket.org/mundipagg/boletoapi/auth"
	"bitbucket.org/mundipagg/boletoapi/models"
)

//Bank é a interface que vai oferecer os serviços em comum entre os bancos
type Bank interface {
	Login() auth.Token
	RegisterBoleto(models.BoletoRequest) models.BoletoResponse
}

var bankRouter map[models.BankNumber]Bank

//InstallBanks instala os bancos configurados no "roteador" de bancos
func InstallBanks() {
	bankRouter = make(map[models.BankNumber]Bank)
	bankRouter[models.BancoDoBrasil] = bankBB{}
}
