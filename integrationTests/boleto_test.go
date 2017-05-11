package integrationTests

import (
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
    "Agreement": {
        "Account": "1231231",
        "AccountDigit": "3",
        "Agency": "1233",
        "AgencyDigit": "2",
        "AgreementNumber": 1014051,
        "Wallet": 17,
        "WalletVariation": 19
    },
    "Authentication": {
        "Password": "eyJpZCI6ImY1NzViYjgtYjBiNy00YSIsImNvZGlnb1B1YmxpY2Fkb3IiOjEwOSwiY29kaWdvU29mdHdhcmUiOjEsInNlcXVlbmNpYWxJbnN0YWxhY2FvIjoxLCJzZXF1ZW5jaWFsQ3JlZGVuY2lhbCI6MX0",
        "Username": "eyJpZCI6IjgwNDNiNTMtZjQ5Mi00YyIsImNvZGlnb1B1YmxpY2Fkb3IiOjEwOSwiY29kaWdvU29mdHdhcmUiOjEsInNlcXVlbmNpYWxJbnN0YWxhY2FvIjoxfQ"
    },
    "BankNumber": 1,
    "Buyer": {
        "Address": {
            "City": "rio de janeiro",
            "Complement": "bloco 3 apto1306",
            "District": "tijuca",
            "Number": "51",
            "StateCode": "rj",
            "Street": "rua morais e silva",
            "ZipCode": "12345678"
        },
        "Document": {
            "Number": "73400584000166",
            "Type": "CNPJ"
        },
        "Name": "moneda"
    },
    "Recipient": {
        "Address": {
            "City": "rio de janeiro",
            "Complement": "2º Piso loja 404",
            "District": "tijuca",
            "Number": "321",
            "StateCode": "rj",
            "Street": "rua da loja",
            "ZipCode": "112312342"
        },
        "Document": {
            "Number": "12312312312366",
            "Type": "CNPJ"
        },
        "Name": "moneda"
    },
    "Title": {
        "AmountInCents": 30000,
        "CreateDate": "0001-01-01T00:00:00Z",
        "ExpireDate": "2017-05-20",
        "ExpireDateTime": "0001-01-01T00:00:00Z",
        "Instructions": "",
        "OurNumber": 101405218
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
	go app.Run(true, true, true)

	Convey("deve-se registrar um boleto e retornar as informações de url, linha digitável e código de barras", t, func() {
		response, st, err := util.Post("http://localhost:3000/v1/boleto/register", getBody(), nil)
		So(err, ShouldEqual, nil)
		So(st, ShouldEqual, 200)
		boleto := models.BoletoResponse{}
		errJSON := json.Unmarshal([]byte(response), &boleto)
		So(errJSON, ShouldEqual, nil)
		Convey("Se o boleto foi registrado então ele tem que está disponível no formato HTML", func() {
			_, st, err := util.Get(boleto.URL, "", nil)
			So(err, ShouldEqual, nil)
			So(st, ShouldEqual, 200)
			//htmlFromBoleto := strings.Contains(html, `<td class="boletoNumber center"><span>`+boleto.DigitableLine+`</span></td>`)
			//So(htmlFromBoleto, ShouldBeTrue)
		})
	})

}

func BenchmarkRegisterBoleto(b *testing.B) {
	go app.Run(true, true, true)
	for i := 0; i < b.N; i++ {
		util.Post("http://localhost:3000/v1/boleto/register", body, nil)

	}
}
