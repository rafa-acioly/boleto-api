package bank

import (
	"fmt"

	"github.com/PMoneda/gonnie"

	"encoding/json"

	"bitbucket.org/mundipagg/boletoapi/config"
	"bitbucket.org/mundipagg/boletoapi/letters"
	"bitbucket.org/mundipagg/boletoapi/log"
	"bitbucket.org/mundipagg/boletoapi/models"
	"bitbucket.org/mundipagg/boletoapi/tmpl"
	"bitbucket.org/mundipagg/boletoapi/util"
)

type bankCaixa struct {
	validate *models.Validator
	log      *log.Log
}

func newCaixa() bankCaixa {
	b := bankCaixa{
		validate: models.NewValidator(),
		log:      log.CreateLog(),
	}
	b.validate.Push(baseValidateAmountInCents)
	b.validate.Push(baseValidateExpireDate)
	b.validate.Push(baseValidateBuyerDocumentNumber)
	b.validate.Push(baseValidateRecipientDocumentNumber)
	b.validate.Push(caixaValidateAccountAndDigit)
	b.validate.Push(caixaValidateAgency)
	return b
}

//Log retorna a referencia do log
func (b bankCaixa) Log() *log.Log {
	return b.log
}

func (b bankCaixa) RegisterBoleto(boleto *models.BoletoRequest) (models.BoletoResponse, error) {
	r := gonnie.NewPipe()
	from := gonnie.Transform(letters.GetResponseTemplateCaixa())
	to := gonnie.Transform(letters.GetRegisterBoletoAPIResponseTmpl())

	bod := r.From("direct:registraBoletoCaixa", boleto)
	bod = bod.To("template://", letters.GetRegisterBoletoCaixaTmpl(), tmpl.GetFuncMaps())
	bod = bod.SetHeader("Content-Type", "text/xml")
	bod = bod.SetHeader("SOAPAction", "IncluiBoleto")
	bod = bod.To(config.Get().URLCaixaRegisterBoleto, map[string]string{"method": "POST", "insecureSkipVerify": "true"})
	ch := bod.Choice()
	ch = ch.When(gonnie.Header("status").IsEqualTo("200"))
	ch = ch.To("transform://?format=xml", from, to, tmpl.GetFuncMaps())

	response := models.BoletoResponse{}
	json.Unmarshal([]byte(bod.GetBody().(string)), &response)
	return response, nil
}
func (b bankCaixa) ProcessBoleto(boleto *models.BoletoRequest) (models.BoletoResponse, error) {
	errs := b.ValidateBoleto(boleto)
	if len(errs) > 0 {
		return models.BoletoResponse{Errors: errs}, nil
	}
	checkSum := b.getCheckSumCode(*boleto)
	boleto.Authentication.AuthorizationToken = b.getAuthToken(checkSum)
	return b.RegisterBoleto(boleto)
}

func (b bankCaixa) ValidateBoleto(boleto *models.BoletoRequest) models.Errors {
	return models.Errors(b.validate.Assert(boleto))
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
