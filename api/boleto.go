package api

import (
	"net/http"

	wkhtmltopdf "github.com/SebastiaanKlippert/go-wkhtmltopdf"

	"errors"

	"strings"

	"bitbucket.org/mundipagg/boletoapi/bank"
	"bitbucket.org/mundipagg/boletoapi/boleto"
	"bitbucket.org/mundipagg/boletoapi/db"
	"bitbucket.org/mundipagg/boletoapi/log"
	"bitbucket.org/mundipagg/boletoapi/models"
	gin "gopkg.in/gin-gonic/gin.v1"
)

//Regista um boleto em um determinado banco
func registerBoleto(c *gin.Context) {
	_boleto, _ := c.Get("boleto")
	boleto := _boleto.(models.BoletoRequest)
	bank, err := bank.Get(boleto.BankNumber)
	if checkError(c, err, log.CreateLog()) {
		return
	}
	lg := bank.Log()
	lg.Operation = "RegisterBoleto"
	lg.NossoNumero = boleto.Title.OurNumber
	lg.Recipient = bank.GetBankNumber().BankName()
	c.Set("log", lg)
	lg.Request(boleto, c.Request.URL.RequestURI(), c.Request.Header)
	mongo, err := db.GetDB()
	if checkError(c, err, lg) {

		return
	}
	resp, errR := bank.ProcessBoleto(boleto)
	if checkError(c, errR, lg) {
		return
	}
	lg.Response(resp, c.Request.URL.RequestURI())
	st := http.StatusOK
	if len(resp.Errors) > 0 {
		st = http.StatusBadRequest
	} else {
		boView := models.NewBoletoView(boleto, resp.BarCodeNumber, resp.DigitableLine)
		resp.URL = boView.EncodeURL()

		mongo.SaveBoleto(boView)
	}
	c.JSON(st, resp)
}

func getBoleto(c *gin.Context) {
	c.Status(200)

	id := c.Query("id")
	fmt := c.Query("fmt")
	mongo, errCon := db.GetDB()
	if checkError(c, errCon, log.CreateLog()) {
		return
	}
	bleto, err := mongo.GetBoletoByID(id)
	if err == nil {
		s := boleto.HTML(bleto)
		if fmt == "html" {
			c.Header("Content-Type", "text/html; charset=utf-8")
			c.Writer.WriteString(s)
		} else {
			c.Header("Content-Type", "application/pdf")
			buf, _ := toPdf(s)
			c.Writer.Write(buf)
		}

	} else {
		checkError(c, errors.New("Boleto não encontrado na base de dados"), log.CreateLog())
	}

}

func toPdf(page string) ([]byte, error) {
	pdfg, err := wkhtmltopdf.NewPDFGenerator()
	if err != nil {
		return nil, err
	}
	pdfg.Dpi.Set(600)
	pdfg.NoCollate.Set(false)
	pdfg.PageSize.Set(wkhtmltopdf.PageSizeA4)
	pdfg.AddPage(wkhtmltopdf.NewPageReader(strings.NewReader(page)))
	err = pdfg.Create()
	if err != nil {
		return nil, err
	}
	return pdfg.Bytes(), nil
}
