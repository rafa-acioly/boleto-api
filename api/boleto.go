package api

import (
	"net/http"

	wkhtmltopdf "github.com/SebastiaanKlippert/go-wkhtmltopdf"

	"errors"

	"strings"

	"encoding/json"
	"os"

	"io/ioutil"

	"bitbucket.org/mundipagg/boletoapi/bank"
	"bitbucket.org/mundipagg/boletoapi/boleto"
	"bitbucket.org/mundipagg/boletoapi/config"
	"bitbucket.org/mundipagg/boletoapi/db"
	"bitbucket.org/mundipagg/boletoapi/log"
	"bitbucket.org/mundipagg/boletoapi/models"
	"bitbucket.org/mundipagg/boletoapi/util"
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
	lg.Request(boleto, c.Request.URL.RequestURI(), util.HeaderToMap(c.Request.Header))
	repo, err := db.GetDB()
	if checkError(c, err, lg) {
		return
	}
	resp, errR := bank.ProcessBoleto(&boleto)

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
		errMongo := repo.SaveBoleto(boView)
		if errMongo != nil {
			saveBoletoJSONFile(boView, lg, errMongo)
		}
	}
	c.JSON(st, resp)
}

func saveBoletoJSONFile(boView models.BoletoView, lg *log.Log, err error) {
	lg.Warn(err.Error(), "I could not save your boleto at Database")
	fd, errOpen := os.Create(config.Get().BoletoJSONFileStore + "/boleto_" + boView.UID + ".json")
	if errOpen != nil {
		lg.Fatal(boView, "[BOLETO_ONLINE_CONTINGENCIA]"+errOpen.Error())
	}
	data, _ := json.Marshal(boView)
	_, errW := fd.Write(data)
	if errW != nil {
		lg.Fatal(boView, "[BOLETO_ONLINE_CONTINGENCIA]"+errW.Error())
	}
	fd.Close()
}

func getBoleto(c *gin.Context) {
	c.Status(200)

	id := c.Query("id")
	format := c.Query("fmt")
	repo, errCon := db.GetDB()
	if checkError(c, errCon, log.CreateLog()) {
		return
	}
	bleto, err := repo.GetBoletoByID(id)
	if err != nil {
		uid := util.Decrypt(id)
		fd, err := os.Open("/boleto_" + uid + ".json")
		if err != nil {
			checkError(c, errors.New("Boleto não encontrado na base de dados"), log.CreateLog())
			return
		}
		data, errR := ioutil.ReadAll(fd)
		if errR != nil {
			checkError(c, errors.New("Boleto não encontrado na base de dados"), log.CreateLog())
			return
		}
		json.Unmarshal(data, &bleto)
		fd.Close()
	}

	s := boleto.HTML(bleto, format)
	if format == "html" {
		c.Header("Content-Type", "text/html; charset=utf-8")
		c.Writer.WriteString(s)
	} else {
		c.Header("Content-Type", "application/pdf")
		buf, _ := toPdf(s)
		c.Writer.Write(buf)
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
