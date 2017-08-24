package integrationTests

import (
	"strings"
	"testing"

	. "github.com/smartystreets/goconvey/convey"

	"encoding/json"

	"time"

	"github.com/mundipagg/boleto-api/app"
	"github.com/mundipagg/boleto-api/models"
	"github.com/mundipagg/boleto-api/util"
)

const body = `{
    "bank": 1,
    "authentication": {
        "username": "user",
        "password": "pass"
    },
    "agreement": {
        "agreement_number": 1014051,
        "wallet_variation": 19,
        "agency":"5797",        
        "account":"6685"
    },
    "title": {
        "expire_date": "2017-05-20",
        "amount": 200,
        "ournumber": 101405190,
        "instructions": "Senhor caixa, após o vencimento",
        "document_number": "123456"
    },
    "buyer": {
        "name": "Mundipagg Tecnologia em Pagamentos",
        "document": {
            "type": "CNPJ",
            "number": "73400584000166"
        },
        "address": {
            "street": "R. Conde de Bonfim",
            "number": "123",
            "complement": "Apto",
            "zipcode": "20520051",
            "city": "Rio de Janeiro",
            "district": "Tijuca",
            "state_code": "RJ"
        }
    },
    "recipient": {
      "name": "Mundipagg Tecnologia em Pagamentos",
        "document": {
            "type": "CNPJ",
            "number": "73400584000166"
        },
        "address": {
            "street": "R. Conde de Bonfim",
            "number": "123",
            "complement": "Apto",
            "zipcode": "20520051",
            "city": "Rio de Janeiro",
            "district": "Tijuca",
            "state_code": "RJ"
        }
    }
}
`

func getBody(bank models.BankNumber, v uint64) string {
	req := models.BoletoRequest{}
	json.Unmarshal([]byte(body), &req)
	req.Title.ExpireDate = time.Now().Format("2006-01-02")
	req.Title.ExpireDateTime = time.Now()
	req.BankNumber = bank
	req.Title.AmountInCents = v
	d, _ := json.Marshal(req)
	return string(d)
}

func getModelBody(bank models.BankNumber, v uint64) models.BoletoRequest {
	str := getBody(bank, v)
	return boletoify(str)
}

func boletoify(str string) models.BoletoRequest {
	bo := models.BoletoRequest{}
	err := json.Unmarshal([]byte(str), &bo)
	if err != nil {
		panic(err)
	}
	return bo
}

func stringify(boleto models.BoletoRequest) string {
	d, _ := json.Marshal(boleto)
	return string(d)
}

func TestRegisterBoletoRequest(t *testing.T) {
	param := app.NewParams()
	param.DevMode = true
	param.DisableLog = true
	param.HTTPOnly = true
	param.MockMode = true
	go app.Run(param)
	time.Sleep(10 * time.Second)
	Convey("deve-se registrar um boleto e retornar as informações de url, linha digitável e código de barras", t, func() {
		req := getBody(models.BancoDoBrasil, 200)
		response, st, err := util.Post("http://localhost:3000/v1/boleto/register", req, nil)
		So(err, ShouldEqual, nil)
		So(st, ShouldEqual, 200)
		boleto := models.BoletoResponse{}
		errJSON := json.Unmarshal([]byte(response), &boleto)
		So(errJSON, ShouldEqual, nil)
		Convey("Se o boleto foi registrado então ele tem que está disponível no formato HTML", func() {
			So(len(boleto.Links), ShouldBeGreaterThan, 0)
			html, st, err := util.Get(boleto.Links[0].Href, "", nil)
			So(err, ShouldEqual, nil)
			So(st, ShouldEqual, 200)
			htmlFromBoleto := strings.Contains(html, boleto.DigitableLine)
			So(htmlFromBoleto, ShouldBeTrue)
		})
	})

	Convey("Deve-se retornar a lista de erros ocorridos durante o registro", t, func() {
		response, st, err := util.Post("http://localhost:3000/v1/boleto/register", getBody(models.BancoDoBrasil, 301), nil)
		So(err, ShouldEqual, nil)
		So(st, ShouldEqual, 400)
		boleto := models.BoletoResponse{}
		errJSON := json.Unmarshal([]byte(response), &boleto)
		So(errJSON, ShouldEqual, nil)
		So(len(boleto.Errors), ShouldBeGreaterThan, 0)
		Convey("Deve-se retornar erro quando passar um Nosso Número inválido", func() {
			m := getModelBody(models.BancoDoBrasil, 200)
			m.Title.OurNumber = 999999999999
			response, st, err := util.Post("http://localhost:3000/v1/boleto/register", stringify(m), nil)
			So(err, ShouldEqual, nil)
			So(st, ShouldEqual, 400)
			boleto := models.BoletoResponse{}
			errJSON := json.Unmarshal([]byte(response), &boleto)
			So(errJSON, ShouldEqual, nil)
			So(len(boleto.Errors), ShouldBeGreaterThan, 0)
			So(boleto.Errors[0].Message, ShouldEqual, "Nosso número inválido")
		})

		Convey("Deve-se tratar o número da conta", func() {
			Convey("O número da conta sempre deve ser passado", func() {
				assert := func(bank models.BankNumber) {
					m := getModelBody(bank, 200)
					m.Agreement.Account = ""
					response, st, err := util.Post("http://localhost:3000/v1/boleto/register", stringify(m), nil)
					So(err, ShouldEqual, nil)
					So(st, ShouldEqual, 400)
					boleto := models.BoletoResponse{}
					errJSON := json.Unmarshal([]byte(response), &boleto)
					So(errJSON, ShouldEqual, nil)
					So(len(boleto.Errors), ShouldBeGreaterThan, 0)
					So(strings.Contains(boleto.Errors[0].Message, "Conta inválida, deve conter até"), ShouldBeTrue)
				}
				assert(models.BancoDoBrasil)
			})

			Convey("O tipo de documento do comprador deve ser CPF ou CNPJ", func() {
				assert := func(bank models.BankNumber) {
					m := getModelBody(bank, 200)
					m.Buyer.Document.Type = "FAIL"
					response, st, err := util.Post("http://localhost:3000/v1/boleto/register", stringify(m), nil)
					So(err, ShouldEqual, nil)
					So(st, ShouldEqual, 400)
					boleto := models.BoletoResponse{}
					errJSON := json.Unmarshal([]byte(response), &boleto)
					So(errJSON, ShouldEqual, nil)
					So(len(boleto.Errors), ShouldBeGreaterThan, 0)
					So(boleto.Errors[0].Message, ShouldEqual, "Tipo de Documento inválido")
				}
				assert(models.BancoDoBrasil)
				assert(models.Caixa)
				assert(models.Citibank)
			})

			Convey("O CPF deve ser válido", func() {
				assert := func(bank models.BankNumber) {
					m := getModelBody(models.BancoDoBrasil, 200)
					m.Buyer.Document.Type = "CPF"
					m.Buyer.Document.Number = "ASDA"
					response, st, err := util.Post("http://localhost:3000/v1/boleto/register", stringify(m), nil)
					So(err, ShouldEqual, nil)
					So(st, ShouldEqual, 400)
					boleto := models.BoletoResponse{}
					errJSON := json.Unmarshal([]byte(response), &boleto)
					So(errJSON, ShouldEqual, nil)
					So(len(boleto.Errors), ShouldBeGreaterThan, 0)
					So(boleto.Errors[0].Message, ShouldEqual, "CPF inválido")
				}
				assert(models.BancoDoBrasil)
				assert(models.Caixa)
				assert(models.Citibank)

			})

		})
	})

	Convey("Quando um boleto não existir na base de dados", t, func() {
		Convey("Deve-se retornar um status 404", func() {
			_, st, err := util.Get("http://localhost:3000/boleto?fmt=html&id=90230843492384", getBody(models.Caixa, 200), nil)
			So(err, ShouldBeNil)
			So(st, ShouldEqual, 404)
		})

		Convey("A mensagem de retorno deverá ser Boleto não encontrado", func() {
			resp, _, err := util.Get("http://localhost:3000/boleto?fmt=html&id=90230843492384", getBody(models.Caixa, 200), nil)
			So(err, ShouldBeNil)
			So(resp, ShouldContainSubstring, "Boleto não encontrado na base de dados")
		})

	})

	Convey("Deve-se registrar um boleto na Caixa", t, func() {
		_, st, err := util.Post("http://localhost:3000/v1/boleto/register", getBody(models.Caixa, 200), nil)
		So(err, ShouldBeNil)
		So(st, ShouldEqual, 200)
		Convey("Deve-se gerar um boleto específico para a Caixa", func() {
			//TODO
		})
	})

	Convey("Deve-se retornar um objeto de erro quando não registra um boleto na Caixa", t, func() {
		response, st, err := util.Post("http://localhost:3000/v1/boleto/register", getBody(models.Caixa, 300), nil)
		So(err, ShouldBeNil)
		So(st, ShouldEqual, 400)
		boleto := models.BoletoResponse{}
		errJSON := json.Unmarshal([]byte(response), &boleto)
		So(errJSON, ShouldEqual, nil)
		So(len(boleto.Errors), ShouldBeGreaterThan, 0)
	})

	Convey("Deve-se retornar um erro quando o campo de instruções tem mais de 40 caracteres", t, func() {
		m := getModelBody(models.Caixa, 200)
		m.Title.Instructions = "Senhor caixa, após o vencimento não aceitar o pagamento"
		response, st, err := util.Post("http://localhost:3000/v1/boleto/register", stringify(m), nil)
		So(err, ShouldBeNil)
		So(st, ShouldEqual, 400)
		boleto := models.BoletoResponse{}
		errJSON := json.Unmarshal([]byte(response), &boleto)
		So(errJSON, ShouldEqual, nil)
		So(len(boleto.Errors), ShouldBeGreaterThan, 0)
		So(boleto.Errors[0].Message, ShouldEqual, "O número máximo permitido para instruções é de 40 caracteres")

	})

	Convey("Quando o serviço da caixa estiver offline", t, func() {
		Convey("Deve-se retornar o status 504", func() {
			resp, st, _ := util.Post("http://localhost:3000/v1/boleto/register", getBody(models.Caixa, 504), nil)
			So(st, ShouldEqual, 504)
			So(strings.Contains(resp, "MP504"), ShouldBeTrue)
		})
	})

}

func BenchmarkRegisterBoleto(b *testing.B) {
	param := app.NewParams()
	param.DevMode = true
	param.DisableLog = true
	param.HTTPOnly = true
	param.MockMode = true
	go app.Run(param)
	for i := 0; i < b.N; i++ {
		util.Post("http://localhost:3000/v1/boleto/register", body, nil)

	}
}
