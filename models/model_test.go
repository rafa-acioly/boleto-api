package models

import (
	"testing"

	"bitbucket.org/mundipagg/boletoapi/test"
	. "github.com/smartystreets/goconvey/convey"
)

func TestDocument(t *testing.T) {
	Convey("Espera que o tipo de documento passado seja um CPF", t, func() {
		document := Document{Number: "13245678901ssa", Type: "CPF"}
		So(document.IsCPF(), ShouldBeTrue)
		document.Type = "cPf"
		So(document.IsCPF(), ShouldBeTrue)
		Convey("Espera que o CPF seja válido", func() {
			err := document.ValidateCPF()
			So(err, ShouldBeNil)
			So(len(document.Number), ShouldEqual, 11)
		})
		Convey("Espera que o CPF seja inválido", func() {
			document.Number = "lasjdlf019239098adjal9390jflsadjf9309jfsl"
			err := document.ValidateCPF()
			So(err, ShouldNotBeNil)
		})
	})
	Convey("Espera que o tipo de documento seja um CNPJ", t, func() {
		document := Document{Number: "12345678901326asdfad", Type: "CNPJ"}
		So(document.IsCNPJ(), ShouldBeTrue)
		document.Type = "cnPj"
		So(document.IsCNPJ(), ShouldBeTrue)
		Convey("Espera que o CNPJ seja válido", func() {
			err := document.ValidateCNPJ()
			So(err, ShouldBeNil)
			So(len(document.Number), ShouldEqual, 14)
		})
		Convey("Espera que o CNPJ seja inválido", func() {
			document.Number = "lasjdlf019239098adjal9390jflsadjf9309jfsl"
			err := document.ValidateCNPJ()
			So(err, ShouldNotBeNil)
		})
	})
}

func TestTitle(t *testing.T) {
	Convey("O DocumentNumber deve conter 10 dígitos", t, func() {
		h := Title{DocumentNumber: "1234567891011"}
		err := h.ValidateDocumentNumber()
		So(err, ShouldBeNil)
		So(len(h.DocumentNumber), ShouldEqual, 10)

		Convey("O DocumentNumber mesmo com menos de 10 dígitos deve possuir 10 dígitos após ser validado com 0 a esquerda", func() {
			h.DocumentNumber = "123x"
			h.ValidateDocumentNumber()
			So(len(h.DocumentNumber), ShouldEqual, 10)
		})

		Convey("O DocumentNumber quando não possuir dígitos deve ser vazio", func() {
			h.DocumentNumber = "xx"
			h.ValidateDocumentNumber()
			So(h.DocumentNumber, ShouldBeEmpty)
		})

		Convey("O DocumentNumber quando for vazio deve permanecer vazio", func() {
			h.DocumentNumber = ""
			h.ValidateDocumentNumber()
			So(h.DocumentNumber, ShouldBeEmpty)
		})
	})

	Convey("As instruções devem ser válidas", t, func() {
		h := Title{Instructions: "Some instructions"}
		err := h.ValidateInstructionsLength(100)
		So(err, ShouldBeNil)
		Convey("As instruções devem ser inválidas", func() {
			err = h.ValidateInstructionsLength(1)
			So(err, ShouldNotBeNil)
		})

	})
}

func TestShouldReturnBankNumberIsValid(t *testing.T) {
	var b BankNumber = 237

	if b.IsBankNumberValid() == false {
		t.Fail()
	}
}

func TestShouldAppendCollectionOfErrrors(t *testing.T) {
	e := NewErrorCollection(ErrorResponse{Code: "200", Message: "Hue2"})
	e.Append("100", "Hue")
	test.ExpectTrue(len(e) == 2, t)
}

func TestShouldCreateNewSingleErrorCollection(t *testing.T) {
	e := NewSingleErrorCollection("200", "Hue2")
	test.ExpectTrue(len(e) == 1, t)
}

func TestIsAgencyValid(t *testing.T) {
	Convey("Deve retornar um erro para a agência inválida", t, func() {
		a := Agreement{
			Agency: "234-2222a",
		}
		err := a.IsAgencyValid()
		So(err, ShouldNotBeNil)

		Convey("Deve ajustar a agência para ter a quantidade certa de dígitos", func() {
			a.Agency = "321"
			err := a.IsAgencyValid()
			So(a.Agency, ShouldEqual, "0321")
			So(err, ShouldBeNil)
		})
	})
}

func TestCalculateAgencyDigit(t *testing.T) {
	Convey("Deve ajustar o dígito da Agência quando ela tiver caracteres inválidos", t, func() {
		a := new(Agreement)
		a.AgencyDigit = "2sssss"
		c := func(s string) string {
			return "1"
		}
		a.CalculateAgencyDigit(c)
		So(a.AgencyDigit, ShouldEqual, "2")
		Convey("Deve calcular o dígito da Agência quando o fornecido for errado", func() {
			a.AgencyDigit = "332sssss"
			a.CalculateAgencyDigit(c)
			So(a.AgencyDigit, ShouldEqual, "1")
		})
	})
}

func TestCalculateAccountDigit(t *testing.T) {
	Convey("Deve ajustar o dígito da Conta quando ela tiver caracteres inválidos", t, func() {
		a := new(Agreement)
		a.AccountDigit = "2sssss"
		c := func(s, y string) string {
			return "1"
		}
		a.CalculateAccountDigit(c)
		So(a.AccountDigit, ShouldEqual, "2")
		Convey("Deve calcular o dígito da Conta quando o fornecido for errado", func() {
			a.AccountDigit = "332sssss"
			a.CalculateAccountDigit(c)
			So(a.AccountDigit, ShouldEqual, "1")
		})
	})
}

func TestIsAccountValid(t *testing.T) {
	Convey("Verifica se a conta é valida e formata para o tamanho correto", t, func() {
		a := Agreement{
			Account: "1234fff",
		}
		err := a.IsAccountValid(8)
		So(err, ShouldBeNil)
		So(a.Account, ShouldEqual, "00001234")
		Convey("Verifica se a conta é valida e retorna um erro", func() {
			a.Account = "654654654654654654654654654564"
			err := a.IsAccountValid(8)
			So(err, ShouldNotBeNil)
		})
	})
}
