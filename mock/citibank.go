package mock

import (
	"errors"
	"io/ioutil"
	"strings"

	gin "gopkg.in/gin-gonic/gin.v1"
)

func registerBoletoCiti(c *gin.Context) {
	sData := `
<soapenv:Envelope xmlns:soapenv="http://schemas.xmlsoap.org/soap/envelope/">
   <soapenv:Header/>
   <soapenv:Body>
      <RegisterBoletoResponse>
         <actionCode>0</actionCode>
         <reasonMessage>OK</reasonMessage>
      </RegisterBoletoResponse>
   </soapenv:Body>
</soapenv:Envelope>
	`

	sDataErr := `
<soapenv:Envelope xmlns:soapenv="http://schemas.xmlsoap.org/soap/envelope/">
   <soapenv:Header/>
   <soapenv:Body>
      <RegisterBoletoResponse>
         <actionCode>99</actionCode>
         <reasonMessage>Erro ao Registrar boleto</reasonMessage>
      </RegisterBoletoResponse>
   </soapenv:Body>
</soapenv:Envelope>
	`
	d, _ := ioutil.ReadAll(c.Request.Body)
	xml := string(d)
	if strings.Contains(xml, "<TitlAmt>504</TitlAmt>") {
		c.AbortWithError(504, errors.New("Teste de Erro"))
	} else if strings.Contains(xml, "<TitlAmt>200</TitlAmt>") {
		c.Data(200, "text/xml", []byte(sData))
	} else {
		c.Data(200, "text/xml", []byte(sDataErr))
	}

}
