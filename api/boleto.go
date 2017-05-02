package api

import (
	"net/http"

	"errors"

	"bitbucket.org/mundipagg/boletoapi/bank"
	"bitbucket.org/mundipagg/boletoapi/boleto"
	"bitbucket.org/mundipagg/boletoapi/db"
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

	resp, errR := bank.ProcessBoleto(boleto)
	if checkError(c, errR) {
		return
	}
	lg.Response(resp, c.Request.URL.RequestURI())
	st := http.StatusOK
	if len(resp.Errors) > 0 {
		st = http.StatusBadRequest
	} else {
		boView := models.NewBoletoView(boleto, resp.BarCodeNumber, resp.DigitableLine)
		resp.URL = boView.EncodeURL()
		db.GetDB().SaveBoleto(boView)
	}
	c.JSON(st, resp)
}

func getBoleto(c *gin.Context) {
	c.Status(200)
	c.Header("Content-Type", "text/html; charset=utf-8")
	id := c.Query("id")
	bleto, err := db.GetDB().GetBoletoByID(id)
	if err == nil {
		boleto.HTML(c.Writer, bleto)
	} else {
		checkError(c, errors.New("Boleto não encontrado na base de dados"))
	}

}
