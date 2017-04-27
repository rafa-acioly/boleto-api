package api

import (
	"net/http"

	"encoding/json"

	"errors"

	"bitbucket.org/mundipagg/boletoapi/bank"
	"bitbucket.org/mundipagg/boletoapi/boleto"
	"bitbucket.org/mundipagg/boletoapi/cache"
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
	var id string
	boView := models.NewBoletoView(boleto, resp.BarCodeNumber, resp.DigitableLine)
	resp.URL, id = boView.EncodeURL()
	cache.Set(id, boView.ToJSON())
	c.JSON(st, resp)
}

func getBoleto(c *gin.Context) {
	c.Status(200)
	c.Header("Content-Type", "text/html; charset=utf-8")
	//format := c.Param("fmt")
	id := c.Query("id")
	data := cache.Get(id)
	if data != nil {
		bleto := models.BoletoView{}
		json.Unmarshal([]byte(data.(string)), &bleto)
		boleto.HTML(c.Writer, bleto)
	} else {
		checkError(c, errors.New("Boleto não encontrado na base de dados"))
	}

}
