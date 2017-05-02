package bank

import (
	"fmt"

	"bitbucket.org/mundipagg/boletoapi/log"
	"bitbucket.org/mundipagg/boletoapi/models"
)

//Bank é a interface que vai oferecer os serviços em comum entre os bancos
type Bank interface {
	ProcessBoleto(models.BoletoRequest) (models.BoletoResponse, error)
	RegisterBoleto(models.BoletoRequest) (models.BoletoResponse, error)
	ValidateBoleto(*models.BoletoRequest) models.Errors
	GetBankNumber() models.BankNumber
	Log() *log.Log
}

//Get retorna estrategia de acordo com o banco ou erro caso o banco não exista
func Get(number models.BankNumber) (Bank, error) {
	switch number {
	case models.BancoDoBrasil:
		return newBB(), nil
	case models.Caixa:
		return newCaixa(), nil
	default:
		return nil, models.ErrorResponse{Code: "MPBankNumber", Message: fmt.Sprintf("Banco %d não existe", number)}
	}
}
