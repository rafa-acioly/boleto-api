package bank

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"bitbucket.org/mundipagg/boletoapi/auth"
	"bitbucket.org/mundipagg/boletoapi/models"
	"bitbucket.org/mundipagg/boletoapi/util"
)

type bankBB struct {
}

func (b bankBB) Login(user, password, body string) (auth.Token, error) {
	client := util.DefaultHTTPClient()
	req, err := http.NewRequest("POST", "https://oauth.desenv.bb.com.br:43000/oauth/token", strings.NewReader(body))
	if err != nil {
		return auth.Token{}, err
	}
	req.SetBasicAuth(user, password)
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("Cache-Control", "no-cache")
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
	if tok.Status != http.StatusOK {
		return tok, errors.New(tok.ErrorDescription)
	}
	return tok, nil
}
func (b bankBB) RegisterBoleto(boleto models.BoletoRequest) (models.BoletoResponse, error) {
	body := "grant_type=client_credentials&scope=cobranca.registro-boletos"
	token, err := b.Login(boleto.Authentication.Username, boleto.Authentication.Password, body)
	if err != nil {
		return models.BoletoResponse{}, err
	}
	fmt.Println(token)
	return models.BoletoResponse{}, nil
}
func (b bankBB) GetBankNumber() models.BankNumber {
	return models.BancoDoBrasil
}
