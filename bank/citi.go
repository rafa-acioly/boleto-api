package bank

import (
	"github.com/PMoneda/gonnie"

	"bitbucket.org/mundipagg/boletoapi/config"
	"bitbucket.org/mundipagg/boletoapi/letters"
	"bitbucket.org/mundipagg/boletoapi/log"
	"bitbucket.org/mundipagg/boletoapi/models"
	"bitbucket.org/mundipagg/boletoapi/tmpl"
	"bitbucket.org/mundipagg/boletoapi/util"
)

type bankCiti struct {
	validate *models.Validator
	log      *log.Log
}

func newCiti() bankCiti {
	b := bankCiti{
		validate: models.NewValidator(),
		log:      log.CreateLog(),
	}
	b.validate.Push(baseValidateAmountInCents)
	b.validate.Push(baseValidateExpireDate)
	b.validate.Push(baseValidateBuyerDocumentNumber)
	b.validate.Push(baseValidateRecipientDocumentNumber)
	return b
}

//Log retorna a referencia do log
func (b bankCiti) Log() *log.Log {
	return b.log
}
func (b bankCiti) RegisterBoleto(boleto *models.BoletoRequest) (models.BoletoResponse, error) {
	r := gonnie.NewPipe()
	serviceURL := config.Get().URLCiti
	from := gonnie.Transform(letters.GetResponseTemplateCiti())
	to := gonnie.Transform(letters.GetRegisterBoletoAPIResponseTmpl())
	bod := r.From("message://?source=inline", boleto, letters.GetRegisterBoletoCitiTmpl(), tmpl.GetFuncMaps())
	bod = bod.To("logseq://?type=request&url="+serviceURL, b.log)
	bod = bod.To(serviceURL, map[string]string{"method": "POST", "insecureSkipVerify": "true"})
	bod = bod.To("logseq://?type=response&url="+serviceURL, b.log)
	ch := bod.Choice()
	ch = ch.When(gonnie.Header("status").IsEqualTo("200"))
	ch = ch.To("transform://?format=xml", from, to, tmpl.GetFuncMaps())
	ch = ch.Otherwise()
	ch = ch.To("logseq://?type=response&url="+serviceURL, b.log).To("apierro://")
	switch t := bod.GetBody().(type) {
	case string:
		response := models.BoletoResponse{}
		util.ParseJSON(t, &response)
		return response, nil
	case models.BoletoResponse:
		return t, nil
	}
	return models.BoletoResponse{}, models.NewInternalServerError("MP500", "Erro interno")
}
func (b bankCiti) ProcessBoleto(boleto *models.BoletoRequest) (models.BoletoResponse, error) {
	errs := b.ValidateBoleto(boleto)
	if len(errs) > 0 {
		return models.BoletoResponse{Errors: errs}, nil
	}
	return b.RegisterBoleto(boleto)
}

func (b bankCiti) ValidateBoleto(boleto *models.BoletoRequest) models.Errors {
	return models.Errors(b.validate.Assert(boleto))
}

//GetBankNumber retorna o codigo do banco
func (b bankCiti) GetBankNumber() models.BankNumber {
	return models.Citibank
}
