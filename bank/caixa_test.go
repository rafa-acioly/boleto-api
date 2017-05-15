package bank

import (
	"encoding/json"
	"fmt"
	"testing"

	"bitbucket.org/mundipagg/boletoapi/letters"
	"bitbucket.org/mundipagg/boletoapi/models"
	"bitbucket.org/mundipagg/boletoapi/tmpl"
	"bitbucket.org/mundipagg/boletoapi/util"

	"time"

	. "github.com/smartystreets/goconvey/convey"
)

const body = `
{

    "BankNumber": 1,

    "Authentication": {

        "Username": "eyJpZCI6IjgwNDNiNTMtZjQ5Mi00YyIsImNvZGlnb1B1YmxpY2Fkb3IiOjEwOSwiY29kaWdvU29mdHdhcmUiOjEsInNlcXVlbmNpYWxJbnN0YWxhY2FvIjoxfQ",

        "Password": "eyJpZCI6IjBjZDFlMGQtN2UyNC00MGQyLWI0YSIsImNvZGlnb1B1YmxpY2Fkb3IiOjEwOSwiY29kaWdvU29mdHdhcmUiOjEsInNlcXVlbmNpYWxJbnN0YWxhY2FvIjoxLCJzZXF1ZW5jaWFsQ3JlZGVuY2lhbCI6MX0"

    },

    "Agreement": {

        "AgreementNumber": 1014051,

        "WalletVariation": 19,

        "Agency":"5797",
        
        "Account":"6685"

    },

    "Title": {

        "ExpireDate": "2017-05-20",

        "AmountInCents": 200,

        "OurNumber": 101405190,

        "Instructions": "Senhor caixa, após o vencimento não aceitar o pagamento",

        "DocumentNumber": "123456"

    },

    "Buyer": {

        "Name": "Mundipagg Tecnologia em Pagamentos",

        "Document": {

            "Type": "CNPJ",

            "Number": "73400584000166"

        },

        "Address": {

            "Street": "R. Conde de Bonfim",

            "Number": "123",

            "Complement": "Apto",

            "ZipCode": "20520051",

            "City": "Rio de Janeiro",

            "District": "Tijuca",

            "StateCode": "RJ"

        }

    },

    "Recipient": {

      "Name": "Mundipagg Tecnologia em Pagamentos",

        "Document": {

            "Type": "CNPJ",

            "Number": "73400584000166"

        },

        "Address": {

            "Street": "R. Conde de Bonfim",

            "Number": "123",

            "Complement": "Apto",

            "ZipCode": "20520051",

            "City": "Rio de Janeiro",

            "District": "Tijuca",

            "StateCode": "RJ"

        }

    }

}
`

func getBody() string {
	req := models.BoletoRequest{}
	json.Unmarshal([]byte(body), &req)
	req.Title.ExpireDate = time.Now().Add(1000 * time.Hour).Format("2006-01-02")
	req.Title.ExpireDateTime = time.Now().Add(1000 * time.Hour)
	d, _ := json.Marshal(req)
	return string(d)
}

func TestGetCaixaCheckSumInfo(t *testing.T) {
	boleto := models.BoletoRequest{
		Agreement: models.Agreement{
			AgreementNumber: 2004001,
		},
		Title: models.Title{
			OurNumber:      1352634,
			ExpireDateTime: time.Date(2017, 5, 20, 12, 12, 12, 12, time.Local),
			AmountInCents:  13567,
		},
		Recipient: models.Recipient{
			Document: models.Document{
				Number: "10497233000103",
			},
		},
	}
	caixa := newCaixa()
	Convey("Geração do token de autorização da Caixa", t, func() {
		Convey("Deve-se formar uma string seguindo o padrão da documentação", func() {
			So(caixa.getCheckSumCode(boleto), ShouldEqual, "2004001000200400113526342005201700000000001356710497233000103")
		})
		Convey("Deve-se fazer um hash sha256 e encodar com base64", func() {
			hash := util.Sha256("2004001000200400113526342005201700000000001356710497233000103")
			So(hash, ShouldEqual, caixa.getAuthToken(caixa.getCheckSumCode(boleto)))
		})
	})

}

func TestBoletoRegisterCaixa(t *testing.T) {
	boleto := models.BoletoRequest{}
	json.Unmarshal([]byte(getBody()), &boleto)
	caixa := newCaixa()
	boleto.Authentication.Password = caixa.getAuthToken(caixa.getCheckSumCode(boleto))
	boleto.Authentication.Username = "Moneda"
	boleto.Title.Instructions = "não aceitar após o vencimento"
	Convey("Deve-se montar a mensagem para o serviço Incluir Boleto", t, func() {
		b := tmpl.New()
		xml, err := b.From(boleto).To(letters.GetRegisterBoletoCaixaTmpl()).XML().Transform()
		So(err, ShouldBeNil)
		fmt.Println()
		fmt.Println(xml)

		resp, _, _ := util.Post("https://des.barramento.caixa.gov.br/sibar/ManutencaoCobrancaBancaria/Boleto/Externo", xml, map[string]string{"Content-Type": "text/xml", "SOAPAction": "IncluiBoleto"})
		fmt.Println(resp)
	})
}
