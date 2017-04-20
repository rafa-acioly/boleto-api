package bank

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"strings"

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
	log *log.Log
}

//Log retorna a referencia do log
func (b bankBB) Log() *log.Log {
	return b.log
}
func (b bankBB) Login(user, password, body string) (auth.Token, error) {
	client := util.DefaultHTTPClient()
	req, err := http.NewRequest("POST", config.Get().URLBBToken, strings.NewReader(body))
	if err != nil {
		return auth.Token{}, err
	}

	req.SetBasicAuth(user, password)
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("Cache-Control", "no-cache")

	b.log.Request(struct {
		Username string
		Password string
		Body     string
	}{
		Username: user,
		Password: password,
		Body:     body,
	}, config.Get().URLBBToken, req.Header)

	resp, errResp := client.Do(req)
	if errResp != nil {
		return auth.Token{}, errResp
	}
	defer resp.Body.Close()
	data, errResponse := ioutil.ReadAll(resp.Body)
	if errResponse != nil {
		return auth.Token{}, errResponse
	}

	tok := auth.Token{Status: resp.StatusCode}
	errParser := json.Unmarshal(data, &tok)
	if errParser != nil {
		return auth.Token{}, errParser
	}

	b.log.Response(tok, config.Get().URLBBToken)

	if tok.Status != http.StatusOK {
		return tok, errors.New(tok.ErrorDescription)
	}
	return tok, nil
}
func (b bankBB) RegisterBoleto(boleto models.BoletoRequest) (string, error) {
	//Body do request necessário para pegar o token do registrar boleto
	body := "grant_type=client_credentials&scope=cobranca.registro-boletos"
	token, err := b.Login(boleto.Authentication.Username, boleto.Authentication.Password, body)
	if err != nil {
		return "", err
	}
	builder := tmpl.New()
	soap, err := builder.From(boleto).To(letters.GetRegisterBoletoBBTmpl()).XML().Transform()

	if err != nil {
		return "", err
	}

	response, status, errRegister := b.registerBoletoRequest(soap, token)

	if errRegister != nil {
		return "", errRegister
	}
	if status != http.StatusOK {
		value, _ := parser.ExtractValues(response, letters.GetRegisterBoletoError())
		j, _ := json.Marshal(
			models.BoletoResponse{
				StatusCode: http.StatusBadRequest,
				Errors:     models.NewSingleErrorCollection(value["faultCode"], value["messageString"]),
			},
		)
		return string(j), nil
	}
	value, _ := parser.ExtractValues(response, letters.GetRegisterBoletoReponseTranslator())
	j, errJSON := builder.From(value).To(letters.GetRegisterBoletoAPIResponseTmpl()).Transform()
	if errJSON != nil {
		return "", errJSON
	}
	return j, nil
}

func (b bankBB) ValidateBoleto(boleto models.BoletoRequest) []string {
	return nil
}

//GetBankNumber retorna o codigo do banco
func (b bankBB) GetBankNumber() models.BankNumber {
	return models.BancoDoBrasil
}

//registerBoletoRequest faz a requisição no serviço do banco para registro de boleto
func (b bankBB) registerBoletoRequest(message string, token auth.Token) (string, int, error) {
	client := util.DefaultHTTPClient()
	body := strings.NewReader(message)
	req, err := http.NewRequest("POST", config.Get().URLBBRegisterBoleto, body)
	if err != nil {
		return "", http.StatusInternalServerError, err
	}
	req.Header.Add("SOAPACTION", "registrarBoleto")
	req.Header.Add("Authorization", "Bearer "+token.AccessToken)
	req.Header.Add("Content-Type", "text/xml; charset=utf-8")

	b.log.Request(message, config.Get().URLBBRegisterBoleto, req.Header)

	resp, errResp := client.Do(req)
	if errResp != nil {
		return "", resp.StatusCode, errResp
	}
	defer resp.Body.Close()
	data, errResponse := ioutil.ReadAll(resp.Body)
	if errResponse != nil {
		return "", resp.StatusCode, errResponse
	}

	sData := string(data)
	b.log.Response(sData, config.Get().URLBBRegisterBoleto)

	return sData, resp.StatusCode, nil
}
