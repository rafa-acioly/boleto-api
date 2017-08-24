package integrationTests

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/mundipagg/boleto-api/app"
	"github.com/mundipagg/boleto-api/models"
	"github.com/mundipagg/boleto-api/util"
	. "github.com/smartystreets/goconvey/convey"
	gin "gopkg.in/gin-gonic/gin.v1"
)

func TestRegisterBoletoCiti(t *testing.T) {
	param := app.NewParams()
	param.DevMode = true
	param.DisableLog = true
	param.HTTPOnly = true
	param.MockMode = true
	router := gin.New()
	app.Run(param, router)

	Convey("Deve-se registrar um boleto no Citi", t, func() {
		boletoReq := getModelBody(models.Citibank, 200)
		boletoReq.Agreement.Account = "088721489"
		boletoReq.Agreement.AccountDigit = "1"
		boletoReq.Agreement.Wallet = 100
		req, err := http.NewRequest("POST", "/v1/boleto/register", strings.NewReader(util.Stringify(boletoReq)))
		w := httptest.NewRecorder()
		router.ServeHTTP(w, req)
		resp := w.Body.String()

		boletoResp := util.ParseJSON(resp, new(models.BoletoResponse)).(*models.BoletoResponse)
		So(err, ShouldBeNil)
		So(w.Code, ShouldEqual, 200)
		savedBoleto := util.ParseJSON(resp, new(models.BoletoView)).(*models.BoletoView)
		So(strings.Contains(boletoResp.Links[0].Href, fmt.Sprintf("%d", savedBoleto.Boleto.Title.OurNumber)), ShouldBeTrue)
		So(strings.Contains(boletoResp.Links[0].Href, savedBoleto.Boleto.Recipient.Document.Number), ShouldBeTrue)
		So(strings.Contains(boletoResp.Links[0].Href, savedBoleto.Boleto.Buyer.Document.Number), ShouldBeTrue)
	})
	Convey("Deve-se validar os campos do agreement", t, func() {
		boletoReq := getModelBody(models.Citibank, 200)
		boletoReq.Agreement.Account = "0887214811111"
		boletoReq.Agreement.AccountDigit = ""
		boletoReq.Agreement.Wallet = 10
		w := httptest.NewRecorder()
		req, err := http.NewRequest("POST", "/v1/boleto/register", strings.NewReader(util.Stringify(boletoReq)))
		router.ServeHTTP(w, req)
		resp := w.Body.String()
		boletoResp := util.ParseJSON(resp, new(models.BoletoResponse)).(*models.BoletoResponse)
		So(boletoResp, ShouldNotBeNil)
		So(err, ShouldBeNil)
		So(w.Code, ShouldEqual, 400)
	})
}
