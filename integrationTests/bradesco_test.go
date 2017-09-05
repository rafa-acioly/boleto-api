package integrationTests

import (
	"strings"
	"testing"

	"github.com/mundipagg/boleto-api/app"
	"github.com/mundipagg/boleto-api/models"
	"github.com/mundipagg/boleto-api/util"
	. "github.com/smartystreets/goconvey/convey"
)

func TestRegisterBoletoBradesco(t *testing.T) {
	param := app.NewParams()
	param.DevMode = true
	param.DisableLog = true
	param.MockMode = true
	go app.Run(param)
	Convey("Deve-se registrar um boleto no Bradesco", t, func() {
		boletoReq := getModelBody(models.Bradesco, 200)
		boletoReq.Authentication.Username = "90000"
		boletoReq.Authentication.Password = "iofWNOeuYv0lilP3uNmzxXYHYFtKyRESMrz-h0_EWVc"
		boletoReq.Agreement.AgreementNumber = 3027577
		boletoReq.Agreement.Wallet = 25
		boletoReq.Agreement.Agency = "3347"
		boletoReq.Agreement.Account = "506541"
		req := util.Stringify(boletoReq)
		resp, st, err := util.Post("http://localhost:3000/v1/boleto/register", req, nil)
		boletoResp := util.ParseJSON(resp, new(models.BoletoResponse)).(*models.BoletoResponse)
		So(err, ShouldBeNil)
		So(st, ShouldEqual, 200)
		savedBoleto := util.ParseJSON(resp, new(models.BoletoView)).(*models.BoletoView)
		So(strings.Contains(boletoResp.Links[0].Href, savedBoleto.Links[0].Href), ShouldBeTrue)
	})
	Convey("Deve-se validar os campos do agreement", t, func() {
		boletoReq := getModelBody(models.Bradesco, 200)
		boletoReq.Agreement.AgreementNumber = 3027577
		boletoReq.Agreement.Wallet = 0
		boletoReq.Agreement.Agency = "3347"
		boletoReq.Agreement.Account = "506541"
		req := util.Stringify(boletoReq)
		resp, st, err := util.Post("http://localhost:3000/v1/boleto/register", req, nil)
		boletoResp := util.ParseJSON(resp, new(models.BoletoView)).(*models.BoletoView)
		So(boletoResp, ShouldNotBeNil)
		So(err, ShouldBeNil)
		So(st, ShouldEqual, 400)
	})
}
