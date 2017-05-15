package bank

import (
	"testing"

	"bitbucket.org/mundipagg/boletoapi/models"
	"bitbucket.org/mundipagg/boletoapi/util"

	"time"

	. "github.com/smartystreets/goconvey/convey"
)

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
