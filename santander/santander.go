package santander

import (
	"errors"
	"fmt"
	"strings"

	"net/http"

	"github.com/PMoneda/flow"
	"github.com/mundipagg/boleto-api/config"
	"github.com/mundipagg/boleto-api/log"
	"github.com/mundipagg/boleto-api/models"
	"github.com/mundipagg/boleto-api/tmpl"
	"github.com/mundipagg/boleto-api/util"
	"github.com/mundipagg/boleto-api/validations"
)

type bankSantander struct {
	validate  *models.Validator
	log       *log.Log
	transport *http.Transport
}

//New Create a new Santander Integration Instance
func New() bankSantander {
	b := bankSantander{
		validate: models.NewValidator(),
		log:      log.CreateLog(),
	}
	b.validate.Push(validations.ValidateAmount)
	b.validate.Push(validations.ValidateExpireDate)
	b.validate.Push(validations.ValidateBuyerDocumentNumber)
	b.validate.Push(validations.ValidateRecipientDocumentNumber)
	t, err := util.BuildTLSTransport(config.Get().SantanderCrtPath, config.Get().SantanderKeyPath, config.Get().SantanderCaPath)
	if err != nil {
		//TODO
	}
	b.transport = t

	return b
}

//Log retorna a referencia do log
func (b bankSantander) Log() *log.Log {
	return b.log
}
func (b bankSantander) GetTicket(boleto *models.BoletoRequest) (string, error) {
	pipe := flow.NewFlow()
	url := config.Get().URLTicketSantander
	tlsUrl := strings.Replace(config.Get().URLTicketSantander, "https", "tls", 1)
	pipe.From("message://?source=inline", boleto, getRequestTicket(), tmpl.GetFuncMaps())
	pipe.To("logseq://?type=request&url="+url, b.log)
	pipe.To(tlsUrl, b.transport)
	pipe.To("logseq://?type=response&url="+url, b.log)
	ch := pipe.Choice()
	ch.When(flow.Header("status").IsEqualTo("200"))
	ch.To("transform://?format=xml", getTicketResponse(), `{{.returnCode}}:::{{unscape .ticket}}:::{{.message}}`, tmpl.GetFuncMaps())
	ch.When(flow.Header("status").IsEqualTo("403"))
	ch.To("print://?msg=Hello").To("set://?prop=body", errors.New("403 Forbidden"))
	ch.Otherwise()
	ch.To("logseq://?type=request&url="+url, b.log).To("set://?prop=body", errors.New("integration error"))
	switch t := pipe.GetBody().(type) {
	case string:
		items := pipe.GetBody().(string)
		parts := strings.Split(items, ":::")
		returnCode, ticket := parts[0], parts[1]
		fmt.Println(items)
		return ticket, checkError(returnCode)
	case error:
		fmt.Println(t.Error())
		return "", t
	}
	return "", nil
}

func (b bankSantander) RegisterBoleto(boleto *models.BoletoRequest) (models.BoletoResponse, error) {
	r := flow.NewFlow()
	serviceURL := config.Get().URLRegisterBoletoSantander
	from := getResponseSantander()
	to := getAPIResponseSantander()
	bod := r.From("message://?source=inline", boleto, getRequestSantander(), tmpl.GetFuncMaps())
	bod = bod.To("logseq://?type=request&url="+serviceURL, b.log)
	bod = bod.To(serviceURL, map[string]string{"method": "POST", "insecureSkipVerify": "true"})
	bod = bod.To("logseq://?type=response&url="+serviceURL, b.log)
	ch := bod.Choice()
	ch = ch.When(flow.Header("status").IsEqualTo("200"))
	ch = ch.To("transform://?format=xml", from, to, tmpl.GetFuncMaps())
	//ch = ch.To("marshall://?format=json", new(models.BoletoResponse))
	ch = ch.Otherwise()
	ch = ch.To("logseq://?type=response&url="+serviceURL, b.log).To("apierro://")
	switch t := bod.GetBody().(type) {
	case string:
		response := util.ParseJSON(t, new(models.BoletoResponse)).(*models.BoletoResponse)
		return *response, nil
	case error:
		fmt.Println(t)
	case models.BoletoResponse:
		return t, nil
	}
	return models.BoletoResponse{}, models.NewInternalServerError("Erro interno", "MP500")
}
func (b bankSantander) ProcessBoleto(boleto *models.BoletoRequest) (models.BoletoResponse, error) {
	errs := b.ValidateBoleto(boleto)
	if len(errs) > 0 {
		return models.BoletoResponse{Errors: errs}, nil
	}
	if ticket, err := b.GetTicket(boleto); err != nil {
		return models.BoletoResponse{Errors: errs}, err
	} else {
		boleto.Authentication.AuthorizationToken = ticket
	}
	return b.RegisterBoleto(boleto)
}

func (b bankSantander) ValidateBoleto(boleto *models.BoletoRequest) models.Errors {
	return models.Errors(b.validate.Assert(boleto))
}

//GetBankNumber retorna o codigo do banco
func (b bankSantander) GetBankNumber() models.BankNumber {
	return models.Santander
}
