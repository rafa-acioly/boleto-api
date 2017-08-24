package integrationTests

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
	gin "gopkg.in/gin-gonic/gin.v1"

	"encoding/json"

	"time"

	"github.com/mundipagg/boleto-api/app"
	"github.com/mundipagg/boleto-api/models"
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
	router := gin.New()
	app.Run(param, router)

	Convey("deve-se registrar um boleto e retornar as informações de url, linha digitável e código de barras", t, func() {
		req, err := http.NewRequest("POST", "/v1/boleto/register", strings.NewReader(getBody(models.BancoDoBrasil, 200)))
		w := httptest.NewRecorder()
		router.ServeHTTP(w, req)
		response := w.Body.String()
		So(err, ShouldEqual, nil)
		So(w.Code, ShouldEqual, 200)
		boleto := models.BoletoResponse{}
		errJSON := json.Unmarshal([]byte(response), &boleto)
		So(errJSON, ShouldEqual, nil)
		w.Flush()
		Convey("Se o boleto foi registrado então ele tem que está disponível no formato HTML", func() {
			So(len(boleto.Links), ShouldBeGreaterThan, 0)
			w := httptest.NewRecorder()
			req, err := http.NewRequest("GET", "/boleto?fmt=html&id="+boleto.ID, strings.NewReader(getBody(models.BancoDoBrasil, 200)))
			router.ServeHTTP(w, req)
			So(err, ShouldEqual, nil)
			So(w.Code, ShouldEqual, 200)
			html := w.Body.String()
			htmlFromBoleto := strings.Contains(html, boleto.DigitableLine)
			So(htmlFromBoleto, ShouldBeTrue)
			w.Flush()
		})
	})

	Convey("Deve-se retornar a lista de erros ocorridos durante o registro", t, func() {
		req, err := http.NewRequest("POST", "/v1/boleto/register", strings.NewReader(getBody(models.BancoDoBrasil, 301)))
		w := httptest.NewRecorder()
		router.ServeHTTP(w, req)
		response := w.Body.String()
		So(err, ShouldEqual, nil)
		So(w.Code, ShouldEqual, 400)
		boleto := models.BoletoResponse{}
		errJSON := json.Unmarshal([]byte(response), &boleto)
		So(errJSON, ShouldEqual, nil)
		So(len(boleto.Errors), ShouldBeGreaterThan, 0)
		w.Flush()
		Convey("Deve-se retornar erro quando passar um Nosso Número inválido", func() {
			m := getModelBody(models.BancoDoBrasil, 200)
			m.Title.OurNumber = 999999999999
			w := httptest.NewRecorder()
			req, err := http.NewRequest("POST", "/v1/boleto/register", strings.NewReader(stringify(m)))
			router.ServeHTTP(w, req)
			response := w.Body.String()
			So(err, ShouldEqual, nil)
			So(w.Code, ShouldEqual, 400)
			boleto := models.BoletoResponse{}
			errJSON := json.Unmarshal([]byte(response), &boleto)
			So(errJSON, ShouldEqual, nil)
			So(len(boleto.Errors), ShouldBeGreaterThan, 0)
			So(boleto.Errors[0].Message, ShouldEqual, "Nosso número inválido")
			w.Flush()
		})

		Convey("Deve-se tratar o número da conta", func() {
			Convey("O número da conta sempre deve ser passado", func() {
				assert := func(bank models.BankNumber) {
					m := getModelBody(bank, 200)
					m.Agreement.Account = ""
					w := httptest.NewRecorder()
					req, err := http.NewRequest("POST", "/v1/boleto/register", strings.NewReader(stringify(m)))
					router.ServeHTTP(w, req)
					response := w.Body.String()
					So(err, ShouldEqual, nil)
					So(w.Code, ShouldEqual, 400)
					boleto := models.BoletoResponse{}
					errJSON := json.Unmarshal([]byte(response), &boleto)
					So(errJSON, ShouldEqual, nil)
					So(len(boleto.Errors), ShouldBeGreaterThan, 0)
					So(strings.Contains(boleto.Errors[0].Message, "Conta inválida, deve conter até"), ShouldBeTrue)
					w.Flush()
				}
				assert(models.BancoDoBrasil)
			})

			Convey("O tipo de documento do comprador deve ser CPF ou CNPJ", func() {
				assert := func(bank models.BankNumber) {
					m := getModelBody(bank, 200)
					m.Buyer.Document.Type = "FAIL"
					w := httptest.NewRecorder()
					req, err := http.NewRequest("POST", "/v1/boleto/register", strings.NewReader(stringify(m)))
					router.ServeHTTP(w, req)
					response := w.Body.String()
					So(err, ShouldEqual, nil)
					So(w.Code, ShouldEqual, 400)
					boleto := models.BoletoResponse{}
					errJSON := json.Unmarshal([]byte(response), &boleto)
					So(errJSON, ShouldEqual, nil)
					So(len(boleto.Errors), ShouldBeGreaterThan, 0)
					So(boleto.Errors[0].Message, ShouldEqual, "Tipo de Documento inválido")
					w.Flush()
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
					w := httptest.NewRecorder()
					req, err := http.NewRequest("POST", "/v1/boleto/register", strings.NewReader(stringify(m)))
					router.ServeHTTP(w, req)
					response := w.Body.String()
					So(err, ShouldEqual, nil)
					So(w.Code, ShouldEqual, 400)
					boleto := models.BoletoResponse{}
					errJSON := json.Unmarshal([]byte(response), &boleto)
					So(errJSON, ShouldEqual, nil)
					So(len(boleto.Errors), ShouldBeGreaterThan, 0)
					So(boleto.Errors[0].Message, ShouldEqual, "CPF inválido")
					w.Flush()
				}
				assert(models.BancoDoBrasil)
				assert(models.Caixa)
				assert(models.Citibank)

			})

		})
	})

	Convey("Quando um boleto não existir na base de dados", t, func() {
		Convey("Deve-se retornar um status 404", func() {
			req, err := http.NewRequest("GET", "/boleto?fmt=html&id=90230843492384", strings.NewReader(getBody(models.Caixa, 200)))
			w := httptest.NewRecorder()
			router.ServeHTTP(w, req)
			So(err, ShouldBeNil)
			So(w.Code, ShouldEqual, 404)
			w.Flush()
		})

		Convey("A mensagem de retorno deverá ser Boleto não encontrado", func() {
			req, err := http.NewRequest("GET", "/boleto?fmt=html&id=90230843492384", strings.NewReader(getBody(models.Caixa, 200)))
			w := httptest.NewRecorder()
			router.ServeHTTP(w, req)
			So(err, ShouldBeNil)
			So(w.Code, ShouldEqual, 404)
			So(w.Body.String(), ShouldContainSubstring, "Boleto não encontrado na base de dados")
			w.Flush()
		})

	})

	Convey("Deve-se registrar um boleto na Caixa", t, func() {
		req, err := http.NewRequest("POST", "/v1/boleto/register", strings.NewReader(getBody(models.Caixa, 200)))
		w := httptest.NewRecorder()
		router.ServeHTTP(w, req)
		So(err, ShouldBeNil)
		So(w.Code, ShouldEqual, 200)
		w.Flush()
		Convey("Deve-se gerar um boleto específico para a Caixa", func() {
			//TODO
		})
	})

	Convey("Deve-se retornar um objeto de erro quando não registra um boleto na Caixa", t, func() {
		req, err := http.NewRequest("POST", "/v1/boleto/register", strings.NewReader(getBody(models.Caixa, 300)))
		w := httptest.NewRecorder()
		router.ServeHTTP(w, req)
		response := w.Body.String()
		So(err, ShouldBeNil)
		So(w.Code, ShouldEqual, 400)
		boleto := models.BoletoResponse{}
		errJSON := json.Unmarshal([]byte(response), &boleto)
		So(errJSON, ShouldEqual, nil)
		So(len(boleto.Errors), ShouldBeGreaterThan, 0)
		w.Flush()
	})

	Convey("Deve-se retornar um erro quando o campo de instruções tem mais de 40 caracteres", t, func() {
		m := getModelBody(models.Caixa, 200)
		m.Title.Instructions = "Senhor caixa, após o vencimento não aceitar o pagamento"
		w := httptest.NewRecorder()
		req, err := http.NewRequest("POST", "/v1/boleto/register", strings.NewReader(stringify(m)))
		router.ServeHTTP(w, req)
		response := w.Body.String()

		So(err, ShouldBeNil)
		So(w.Code, ShouldEqual, 400)
		boleto := models.BoletoResponse{}
		errJSON := json.Unmarshal([]byte(response), &boleto)
		So(errJSON, ShouldEqual, nil)
		So(len(boleto.Errors), ShouldBeGreaterThan, 0)
		So(boleto.Errors[0].Message, ShouldEqual, "O número máximo permitido para instruções é de 40 caracteres")
		w.Flush()
	})

	Convey("Quando o serviço da caixa estiver offline", t, func() {
		Convey("Deve-se retornar o status 504", func() {
			req, _ := http.NewRequest("POST", "/v1/boleto/register", strings.NewReader(getBody(models.Caixa, 504)))
			w := httptest.NewRecorder()
			router.ServeHTTP(w, req)
			resp := w.Body.String()
			So(w.Code, ShouldEqual, 504)
			So(strings.Contains(resp, "MP504"), ShouldBeTrue)
			w.Flush()
		})
	})

}
