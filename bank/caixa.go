package bank

import (
	"fmt"

	"bitbucket.org/mundipagg/boletoapi/auth"
	"bitbucket.org/mundipagg/boletoapi/log"
	"bitbucket.org/mundipagg/boletoapi/models"
	"bitbucket.org/mundipagg/boletoapi/util"
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

//getCheckSumCode Código do Cedente (7 posições) + Nosso Número (17 posições) + Data de Vencimento (DDMMAAAA) + Valor (15 posições) + CPF/CNPJ (14 Posições)
func (b bankCaixa) getCheckSumCode(boleto models.BoletoRequest) string {
	ourNumber := fmt.Sprintf("%d%d", boleto.Agreement.AgreementNumber, boleto.Title.OurNumber)
	return fmt.Sprintf("%07d%017s%s%015d%014s",
		boleto.Agreement.AgreementNumber,
		ourNumber,
		boleto.Title.ExpireDateTime.Format("02012006"),
		boleto.Title.AmountInCents,
		boleto.Recipient.Document.Number)
}

func (b bankCaixa) getAuthToken(info string) string {
	return util.Sha256(info)
}

//GetBankNumber retorna o codigo do banco
func (b bankCaixa) GetBankNumber() models.BankNumber {
	return models.Caixa
}
