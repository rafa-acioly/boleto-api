package bank

import (
	"encoding/json"
	"errors"
	"net/http"

	"bitbucket.org/mundipagg/boletoapi/auth"
	"bitbucket.org/mundipagg/boletoapi/config"
	"bitbucket.org/mundipagg/boletoapi/letters"
	"bitbucket.org/mundipagg/boletoapi/log"
	"bitbucket.org/mundipagg/boletoapi/models"
	"bitbucket.org/mundipagg/boletoapi/parser"
	"bitbucket.org/mundipagg/boletoapi/tmpl"
	"bitbucket.org/mundipagg/boletoapi/util"
)

type bankBB struct {
	validate *models.Validator
	log      *log.Log
	token    auth.Token
}

//Cria uma nova instância do objeto que implementa os serviços do Banco do Brasil e configura os validadores que serão utilizados
func newBB() bankBB {
	b := bankBB{
		validate: models.NewValidator(),
		log:      log.CreateLog(),
	}
	b.validate.Push(bbValidateAccountAndDigit)
	b.validate.Push(bbValidateAgencyAndDigit)
	b.validate.Push(bbValidateOurNumber)
	b.validate.Push(bbValidateWalletVariation)
	b.validate.Push(baseValidateAmountInCents)
	b.validate.Push(baseValidateExpireDate)
	b.validate.Push(baseValidateBuyerDocumentNumber)
	b.validate.Push(baseValidateRecipientDocumentNumber)
	b.validate.Push(bbValidateTitleInstructions)
	b.validate.Push(bbValidateTitleDocumentNumber)
	return b
}

//Log retorna a referencia do log
func (b bankBB) Log() *log.Log {
	return b.log
}
func (b *bankBB) login(user, password string) (auth.Token, error) {

	body := "grant_type=client_credentials&scope=cobranca.registro-boletos"
	header := make(map[string]string)
	header["Content-Type"] = "application/x-www-form-urlencoded"
	header["Cache-Control"] = "no-cache"
	header["Authorization"] = "Basic " + util.Base64(user+":"+password)
	resp, st, err := util.Post(config.Get().URLBBToken, body, header)
	if err != nil {
		return auth.Token{}, err
	}
	b.log.Request(struct {
		Username string
		Password string
		Body     string
	}{
		Username: user,
		Password: password,
		Body:     body,
	}, config.Get().URLBBToken, header)

	tok := auth.Token{Status: st}
	errParser := json.Unmarshal([]byte(resp), &tok)
	if errParser != nil {
		return auth.Token{}, errParser
	}
	b.log.Response(tok, config.Get().URLBBToken)
	if tok.Status != http.StatusOK {
		return tok, errors.New(tok.ErrorDescription)
	}
	b.token = tok
	return tok, nil
}

//ProcessBoleto faz o processamento de registro de boleto
func (b bankBB) ProcessBoleto(boleto *models.BoletoRequest) (models.BoletoResponse, error) {
	errs := b.ValidateBoleto(boleto)
	if len(errs) > 0 {
		return models.BoletoResponse{Errors: errs}, nil
	}
	_, err := b.login(boleto.Authentication.Username, boleto.Authentication.Password)
	if err != nil {
		return models.BoletoResponse{}, models.NewErrorResponse("MP500", err.Error())
	}
	return b.RegisterBoleto(boleto)
}

func (b bankBB) RegisterBoleto(boleto *models.BoletoRequest) (models.BoletoResponse, error) {
	builder := tmpl.New()
	soap, err := builder.From(boleto).To(letters.GetRegisterBoletoBBTmpl()).XML().Transform()
	if err != nil {
		return models.BoletoResponse{}, err
	}
	response, status, errRegister := b.doRequest(soap, b.token)
	if errRegister != nil {
		return models.BoletoResponse{}, errRegister
	}
	if status != http.StatusOK {
		value, _ := parser.ExtractValues(response, letters.GetRegisterBoletoError())
		j := models.BoletoResponse{
			StatusCode: http.StatusBadRequest,
			Errors:     models.NewSingleErrorCollection(value["faultCode"], value["messageString"]),
		}
		return j, nil
	}

	value, _ := parser.ExtractValues(response, letters.GetRegisterBoletoReponseTranslator())
	j, errJSON := builder.From(value).To(letters.GetRegisterBoletoAPIResponseTmpl()).Transform()
	if errJSON != nil {
		return models.BoletoResponse{}, errJSON
	}
	resp := models.BoletoResponse{}
	errParse := json.Unmarshal([]byte(j), &resp)
	if errParse != nil {
		return models.BoletoResponse{}, errParse
	}
	return resp, nil
}

func (b bankBB) ValidateBoleto(boleto *models.BoletoRequest) models.Errors {
	return models.Errors(b.validate.Assert(boleto))
}

//GetBankNumber retorna o codigo do banco
func (b bankBB) GetBankNumber() models.BankNumber {
	return models.BancoDoBrasil
}

//registerBoletoRequest faz a requisição no serviço do banco para registro de boleto
func (b bankBB) doRequest(message string, token auth.Token) (string, int, error) {
	header := make(map[string]string)
	header["SOAPACTION"] = "registrarBoleto"
	header["Authorization"] = "Bearer " + token.AccessToken
	header["Content-Type"] = "text/xml; charset=utf-8"
	b.log.Request(message, config.Get().URLBBRegisterBoleto, header)
	resp, status, err := util.Post(config.Get().URLBBRegisterBoleto, message, header)
	b.log.Response(resp, config.Get().URLBBRegisterBoleto)
	return resp, status, err

}
