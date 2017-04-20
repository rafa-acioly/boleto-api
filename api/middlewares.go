package api

import (
	"time"

	"bitbucket.org/mundipagg/boletoapi/models"
	gin "gopkg.in/gin-gonic/gin.v1"
)

// ReturnHeaders 'seta' os headers padrões de resposta
func ReturnHeaders() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("Content-Type", "application/json")
		c.Next()
	}
}

//ParseBoleto trata a entrada de boleto em todos os requests
func ParseBoleto() gin.HandlerFunc {
	return func(c *gin.Context) {
		boleto := models.BoletoRequest{}
		errBind := c.BindJSON(&boleto)
		if checkError(c, errBind) {
			return
		}
		d, errFmt := time.Parse("2006-01-02", boleto.Title.ExpireDate)
		boleto.Title.ExpireDateTime = d
		if checkError(c, errFmt) {
			return
		}
		c.Set("boleto", boleto)
		c.Next()
	}
}
