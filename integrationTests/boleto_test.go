package integrationTests

import (
	"strings"
	"testing"

	. "github.com/smartystreets/goconvey/convey"

	"encoding/json"

	"time"

	"bitbucket.org/mundipagg/boletoapi/app"
	"bitbucket.org/mundipagg/boletoapi/models"
	"bitbucket.org/mundipagg/boletoapi/util"
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
	req.Title.ExpireDate = time.Now().Format("2006-01-02")
	d, _ := json.Marshal(req)
	return string(d)
}

func TestRegisterBoletoRequest(t *testing.T) {
	go app.Run(true, true, false)

	Convey("deve-se registrar um boleto e retornar as informações de url, linha digitável e código de barras", t, func() {
		response, st, err := util.Post("http://localhost:3000/v1/boleto/register", getBody(), nil)
		So(err, ShouldEqual, nil)
		So(st, ShouldEqual, 200)
		boleto := models.BoletoResponse{}
		errJSON := json.Unmarshal([]byte(response), &boleto)
		So(errJSON, ShouldEqual, nil)
		Convey("Se o boleto foi registrado então ele tem que está disponível no formato HTML", func() {
			html, st, err := util.Get(boleto.URL, "", nil)
			So(err, ShouldEqual, nil)
			So(st, ShouldEqual, 200)
			htmlFromBoleto := strings.Contains(html, boleto.DigitableLine)
			So(htmlFromBoleto, ShouldBeTrue)
		})
	})

}

func BenchmarkRegisterBoleto(b *testing.B) {
	go app.Run(true, true, true)
	for i := 0; i < b.N; i++ {
		util.Post("http://localhost:3000/v1/boleto/register", body, nil)

	}
}
