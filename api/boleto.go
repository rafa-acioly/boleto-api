package api

import (
	"net/http"

	"time"

	"bitbucket.org/mundipagg/boletoapi/bank"
	"bitbucket.org/mundipagg/boletoapi/boleto"
	"bitbucket.org/mundipagg/boletoapi/models"
	gin "gopkg.in/gin-gonic/gin.v1"
)

//Regista um boleto em um determinado banco
func registerBoleto(c *gin.Context) {
	_boleto, _ := c.Get("boleto")
	boleto := _boleto.(models.BoletoRequest)
	bank, err := bank.Get(boleto.BankNumber)
	if checkError(c, err) {
		return
	}
	lg := bank.Log()
	lg.Operation = "RegisterBoleto"
	lg.NossoNumero = boleto.Title.OurNumber

	lg.Recipient = bank.GetBankNumber().BankName()
	lg.Request(boleto, c.Request.URL.RequestURI(), c.Request.Header)

	resp, errR := bank.RegisterBoleto(boleto)
	if checkError(c, errR) {
		return
	}
	lg.Response(resp, c.Request.URL.RequestURI())
	st := http.StatusOK
	if len(resp.Errors) > 0 {
		st = http.StatusBadRequest
	}
	c.JSON(st, resp)
}

func getBoleto(c *gin.Context) {
	c.Status(200)
	c.Header("Content-Type", "text/html")
	bleto := models.BoletoRequest{
		Title: models.Title{
			ExpireDateTime: time.Now(),
			AmountInCents:  1000,
			OurNumber:      12345678,
		},
		Agreement: models.Agreement{
			Account:      "12345",
			AccountDigit: "1",
			Agency:       "1234",
			Wallet:       157,
		},
		Buyer: models.Buyer{
			Address: models.Address{
				Street:     "Avenida do Pagador",
				Number:     "100",
				Complement: "Bloco 3 Apto 1345",
				City:       "Rio de Janeiro",
				District:   "Tijuca",
				StateCode:  "RJ",
				ZipCode:    "123131313123",
			},
			Document: models.Document{
				Type:   "CPF",
				Number: "12345678923",
			},
			Name: "Nome do Pagador",
		},
		Recipient: models.Recipient{
			Name: "Fulano de Tal Dono da Loja",
			Document: models.Document{
				Type:   "CNPJ",
				Number: "123547568678",
			},
			Address: models.Address{
				Street:     "Rua da Loja vendedora",
				City:       "Sao Paulo",
				Number:     "131",
				Complement: "3º andar sala 3b",
				District:   "Vila Mariana",
				ZipCode:    "1233912839",
				StateCode:  "SP",
			},
		},
		BankNumber: 1,
	}

	boleto.HTML(c.Writer, bleto)
}
