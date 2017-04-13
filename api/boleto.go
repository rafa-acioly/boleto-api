package api

import (
	"net/http"

	"time"

	"bitbucket.org/mundipagg/boletoapi/bank"
	"bitbucket.org/mundipagg/boletoapi/models"
	gin "gopkg.in/gin-gonic/gin.v1"
)

//Regista um boleto em um determinado banco
func registerBoleto(c *gin.Context) {
	boleto := models.BoletoRequest{}
	errBind := c.BindJSON(&boleto)
	//TODO melhorar isso
	d, errFmt := time.Parse("2006-01-02", boleto.Title.ExpireDate)
	boleto.Title.ExpireDateTime = d
	if errFmt != nil {
		c.JSON(http.StatusBadRequest, errorResponse{Code: "000", Message: errFmt.Error()})
		return
	}
	if errBind != nil {
		c.JSON(http.StatusBadRequest, errorResponse{Code: "000", Message: errBind.Error()})
		return
	}
	bank, err := bank.Get(boleto.BankNumber)
	if err != nil {
		c.JSON(http.StatusBadRequest, errorResponse{Code: "001", Message: err.Error()})
		return
	}
	resp, errR := bank.RegisterBoleto(boleto)
	if errR != nil {
		c.Data(http.StatusBadRequest, "application/json", []byte(resp))
		return
	}
	c.Data(http.StatusOK, "application/json", []byte(resp))
}
