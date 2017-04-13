package api

import (
	"net/http"

	"bitbucket.org/mundipagg/boletoapi/bank"
	"bitbucket.org/mundipagg/boletoapi/models"
	gin "gopkg.in/gin-gonic/gin.v1"
)

//Regista um boleto em um determinado banco
func registerBoleto(c *gin.Context) {
	boleto := models.BoletoRequest{}
	c.BindJSON(&boleto)
	bank, err := bank.Get(boleto.BankNumber)
	if err != nil {
		c.JSON(http.StatusBadRequest, errorResponse{Code: "001", Message: err.Error()})
		return
	}
	bank.RegisterBoleto(boleto)
	c.JSON(200, boleto)
}
