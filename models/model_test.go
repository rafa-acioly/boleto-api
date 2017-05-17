package models

import (
	"testing"

	"bitbucket.org/mundipagg/boletoapi/test"
	. "github.com/smartystreets/goconvey/convey"
)

func TestShouldReturnValidCpfOnDocumentType(t *testing.T) {
	document := Document{Number: "12345678901", Type: "CPF"}
	if document.IsCPF() == false {
		t.Fail()
	}
}

func TestShouldValidateDocumentNumber(t *testing.T) {
	h := Title{DocumentNumber: "1234567891011"}
	h.ValidateDocumentNumber()
	test.ExpectTrue(len(h.DocumentNumber) == 10, t)

	h.DocumentNumber = "123x"
	h.ValidateDocumentNumber()
	test.ExpectTrue(len(h.DocumentNumber) == 10, t)

	h.DocumentNumber = "xx"
	h.ValidateDocumentNumber()
	test.ExpectTrue(h.DocumentNumber == "", t)
}

func TestShouldReturnInvalidCpfOnDocumentType(t *testing.T) {
	document := Document{Number: "1234567890132", Type: "CNPJ"}
	if document.IsCNPJ() == false {
		t.Fail()
	}
}

func TestShouldReturnValidCnpjOnDocumentType(t *testing.T) {
	document := Document{Number: "1234567890132", Type: "CnpJ"}
	if document.IsCNPJ() == false {
		t.Fail()
	}
}

func TestShouldReturnInvalidCnpjOnDocumentType(t *testing.T) {
	document := Document{Number: "12345678901", Type: "CPF"}
	if document.IsCNPJ() {
		t.Fail()
	}
}

func TestShouldReturnValidCnpjOnDocumentNumber(t *testing.T) {
	document := Document{Number: "12345678901564fas", Type: "CNPJ"}
	if err := document.ValidateCNPJ(); err != nil {
		t.Fail()
	}
}

func TestShouldReturnInvalidCnpjOnDocumentNumber(t *testing.T) {
	document := Document{Number: "12345678901564asdf22", Type: "CNPJ"}
	if err := document.ValidateCNPJ(); err == nil {
		t.Fail()
	}
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
