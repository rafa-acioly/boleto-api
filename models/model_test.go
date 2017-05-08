package models

import (
	"testing"

	"bitbucket.org/mundipagg/boletoapi/test"
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
	a := Agreement{
		Agency: "234-2a",
	}
	err := a.IsAgencyValid()
	test.ExpectNoError(err, t)
	test.ExpectTrue(a.Agency == "2342", t)
}

func TestIsAgencyInValid(t *testing.T) {
	a := Agreement{
		Agency: "234-2222a",
	}
	err := a.IsAgencyValid()
	test.ExpectError(err, t)
}

func TestIsAgencyValidWithLessThan4Digits(t *testing.T) {
	a := Agreement{
		Agency: "321",
	}
	err := a.IsAgencyValid()
	test.ExpectNoError(err, t)
	test.ExpectTrue(a.Agency == "0321", t)
}

func TestCalculateAgencyDigit(t *testing.T) {
	a := new(Agreement)
	a.AccountDigit = "1sssss"
	c := func(s string) string {
		return "1"
	}
	a.CalculateAgencyDigit(c)
	test.ExpectTrue(a.AgencyDigit == "1", t)
}

func WTestCalculateAgencyDigitWithInvalidDigit(t *testing.T) {
	a := Agreement{
		AgencyDigit: "",
	}
	c := func(s string) string {
		return "1"
	}
	a.CalculateAgencyDigit(c)
	test.ExpectTrue(a.AgencyDigit == "1", t)
}

func TestIsAccountValid(t *testing.T) {
	a := Agreement{
		Account: "1234fff",
	}
	err := a.IsAccountValid(8)
	test.ExpectNoError(err, t)
	test.ExpectTrue(a.Account == "00001234", t)
}

func TestIsAccountInValid(t *testing.T) {
	a := Agreement{
		Account: "123456789",
	}
	err := a.IsAccountValid(8)
	test.ExpectError(err, t)
}
