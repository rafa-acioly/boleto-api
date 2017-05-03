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
	return b
}

//Log retorna a referencia do log
func (b bankBB) Log() *log.Log {
	return b.log
}
func (b *bankBB) login(user, password string) (auth.Token, error) {
	if config.Get().MockMode {
		return auth.Token{AccessToken: "1111111111"}, nil
	}
	body := "grant_type=client_credentials&scope=cobranca.registro-boletos"
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
	b.token = tok
	return tok, nil
}

//ProcessBoleto faz o processamento de registro de boleto
func (b bankBB) ProcessBoleto(boleto models.BoletoRequest) (models.BoletoResponse, error) {
	errs := b.ValidateBoleto(&boleto)
	if len(errs) > 0 {
		return models.BoletoResponse{Errors: errs}, nil
	}
	_, err := b.login(boleto.Authentication.Username, boleto.Authentication.Password)
	if err != nil {
		return models.BoletoResponse{}, models.NewErrorResponse("MP500", err.Error())
	}
	return b.RegisterBoleto(boleto)
}

func (b bankBB) RegisterBoleto(boleto models.BoletoRequest) (models.BoletoResponse, error) {
	builder := tmpl.New()
	soap, err := builder.From(boleto).To(letters.GetRegisterBoletoBBTmpl()).XML().Transform()
	if err != nil {
		return models.BoletoResponse{}, err
	}
	response, status, errRegister := b.registerBoletoRequest(soap, b.token)
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
func (b bankBB) registerBoletoRequest(message string, token auth.Token) (string, int, error) {
	if config.Get().MockMode {
		return b.doMockSuccess(message, token)
	}
	return b.doRequest(message, token)
}

//registerBoletoRequest faz a requisição no serviço do banco para registro de boleto
func (b bankBB) doRequest(message string, token auth.Token) (string, int, error) {
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

func (b bankBB) doMockSuccess(message string, token auth.Token) (string, int, error) {
	errMock := `
		<?xml version="1.0" encoding="UTF-8"?>
		<SOAP-ENV:Envelope xmlns:SOAP-ENV="http://schemas.xmlsoap.org/soap/envelope/">
		<SOAP-ENV:Body>
			<ns0:resposta xmlns:ns0="http://www.tibco.com/schemas/bws_registro_cbr/Recursos/XSD/Schema.xsd">
				<ns0:siglaSistemaMensagem />
				<ns0:codigoRetornoPrograma>5</ns0:codigoRetornoPrograma>
				<ns0:nomeProgramaErro>CBRSR005</ns0:nomeProgramaErro>
				<ns0:textoMensagemErro>?CPF do pagador nao encontrado na base.</ns0:textoMensagemErro>
				<ns0:numeroPosicaoErroPrograma>5</ns0:numeroPosicaoErroPrograma>
				<ns0:codigoTipoRetornoPrograma>0</ns0:codigoTipoRetornoPrograma>
				<ns0:textoNumeroTituloCobrancaBb />
				<ns0:numeroCarteiraCobranca>17</ns0:numeroCarteiraCobranca>
				<ns0:numeroVariacaoCarteiraCobranca>19</ns0:numeroVariacaoCarteiraCobranca>
				<ns0:codigoPrefixoDependenciaBeneficiario>3851</ns0:codigoPrefixoDependenciaBeneficiario>
				<ns0:numeroContaCorrenteBeneficiario>8570</ns0:numeroContaCorrenteBeneficiario>
				<ns0:codigoCliente>932131545</ns0:codigoCliente>
				<ns0:linhaDigitavel />
				<ns0:codigoBarraNumerico />
				<ns0:codigoTipoEnderecoBeneficiario>0</ns0:codigoTipoEnderecoBeneficiario>
				<ns0:nomeLogradouroBeneficiario />
				<ns0:nomeBairroBeneficiario />
				<ns0:nomeMunicipioBeneficiario />
				<ns0:codigoMunicipioBeneficiario>0</ns0:codigoMunicipioBeneficiario>
				<ns0:siglaUfBeneficiario />
				<ns0:codigoCepBeneficiario>0</ns0:codigoCepBeneficiario>
				<ns0:indicadorComprovacaoBeneficiario />
				<ns0:numeroContratoCobranca>17414296</ns0:numeroContratoCobranca>
			</ns0:resposta>
		</SOAP-ENV:Body>
		</SOAP-ENV:Envelope>
	`

	sData := `
		<?xml version="1.0" encoding="UTF-8"?>
		<SOAP-ENV:Envelope xmlns:SOAP-ENV="http://schemas.xmlsoap.org/soap/envelope/">
		<SOAP-ENV:Body>
			<ns0:resposta xmlns:ns0="http://www.tibco.com/schemas/bws_registro_cbr/Recursos/XSD/Schema.xsd">
				<ns0:siglaSistemaMensagem />
				<ns0:codigoRetornoPrograma>0</ns0:codigoRetornoPrograma>
				<ns0:nomeProgramaErro />
				<ns0:textoMensagemErro />
				<ns0:numeroPosicaoErroPrograma>0</ns0:numeroPosicaoErroPrograma>
				<ns0:codigoTipoRetornoPrograma>0</ns0:codigoTipoRetornoPrograma>
				<ns0:textoNumeroTituloCobrancaBb>00010140510000066673</ns0:textoNumeroTituloCobrancaBb>
				<ns0:numeroCarteiraCobranca>17</ns0:numeroCarteiraCobranca>
				<ns0:numeroVariacaoCarteiraCobranca>19</ns0:numeroVariacaoCarteiraCobranca>
				<ns0:codigoPrefixoDependenciaBeneficiario>3851</ns0:codigoPrefixoDependenciaBeneficiario>
				<ns0:numeroContaCorrenteBeneficiario>8570</ns0:numeroContaCorrenteBeneficiario>
				<ns0:codigoCliente>932131545</ns0:codigoCliente>
				<ns0:linhaDigitavel>00190000090101405100500066673179971340000010000</ns0:linhaDigitavel>
				<ns0:codigoBarraNumerico>00199713400000100000000001014051000006667317</ns0:codigoBarraNumerico>
				<ns0:codigoTipoEnderecoBeneficiario>0</ns0:codigoTipoEnderecoBeneficiario>
				<ns0:nomeLogradouroBeneficiario>Cliente nao informado.</ns0:nomeLogradouroBeneficiario>
				<ns0:nomeBairroBeneficiario />
				<ns0:nomeMunicipioBeneficiario />
				<ns0:codigoMunicipioBeneficiario>0</ns0:codigoMunicipioBeneficiario>
				<ns0:siglaUfBeneficiario />
				<ns0:codigoCepBeneficiario>0</ns0:codigoCepBeneficiario>
				<ns0:indicadorComprovacaoBeneficiario />
				<ns0:numeroContratoCobranca>17414296</ns0:numeroContratoCobranca>
			</ns0:resposta>
		</SOAP-ENV:Body>
		</SOAP-ENV:Envelope>
	`
	s := sData
	if strings.Contains(message, "<sch:valorOriginalTitulo>100</sch:valorOriginalTitulo>") {
		s = errMock
	}
	return s, 200, nil
}
