package integrationTests

import (
	"fmt"
	"testing"

	"encoding/json"

	"bitbucket.org/mundipagg/boletoapi/bank"
	"bitbucket.org/mundipagg/boletoapi/letters"
	"bitbucket.org/mundipagg/boletoapi/models"
	"bitbucket.org/mundipagg/boletoapi/tmpl"
	. "github.com/smartystreets/goconvey/convey"
)

func TestCaixaIntegration(t *testing.T) {
	boleto := models.BoletoRequest{}
	json.Unmarshal([]byte(getBody()), &boleto)
	Convey("Quando registrar boleto na Caixa", t, func() {
		Convey("Deve-se montar a string de autenticação", func() {
			caixa, _ := bank.Get(models.Caixa)
			caixa.GetAuthToken(boleto)

		})

		Convey("Deve-se montar a mensagem para o serviço Incluir Boleto", func() {
			b := tmpl.New()

			xml, err := b.From(boleto).To(letters.GetRegisterBoletoCaixaTmpl()).XML().Transform()
			So(err, ShouldBeNil)
			fmt.Println()
			fmt.Println(xml)
		})
	})
}
