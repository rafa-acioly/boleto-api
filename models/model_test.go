package models

import (
	"testing"
	"time"

	"bitbucket.org/mundipagg/boletoapi/test"
)

func TestShouldReturnValidCpfOnDocumentType(t *testing.T) {
	document := Document{Number: "12345678901", Type: "CPF"}
	if document.Type.IsCpf() == false {
		t.Fail()
	}
}

func TestShouldReturnInvalidCpfOnDocumentType(t *testing.T) {
	document := Document{Number: "1234567890132", Type: "CNPJ"}
	if document.Type.IsCpf() {
		t.Fail()
	}
}

func TestShouldReturnValidCnpjOnDocumentType(t *testing.T) {
	document := Document{Number: "1234567890132", Type: "CnpJ"}
	if document.Type.IsCnpj() == false {
		t.Fail()
	}
}

func TestShouldReturnInvalidCnpjOnDocumentType(t *testing.T) {
	document := Document{Number: "12345678901", Type: "CPF"}
	if document.Type.IsCnpj() {
		t.Fail()
	}
}

func TestShouldReturnValidCpfOnDocumentNumber(t *testing.T) {
	document := Document{Number: "12345678901", Type: "CPF"}
	if document.Number.IsCpf() == false {
		t.Fail()
	}
}

func TestShouldReturnInvalidCpfOnDocumentNumber(t *testing.T) {
	document := Document{Number: "12345678901asd", Type: "CPF"}
	if document.Number.IsCpf() {
		t.Fail()
	}
}

func TestShouldReturnValidCnpjOnDocumentNumber(t *testing.T) {
	document := Document{Number: "12345678901564", Type: "CNPJ"}
	if document.Number.IsCnpj() == false {
		t.Fail()
	}
}

func TestShouldReturnInvalidCnpjOnDocumentNumber(t *testing.T) {
	document := Document{Number: "12345678901564asdf", Type: "CNPJ"}
	if document.Number.IsCnpj() {
		t.Fail()
	}
}

func TestShouldReturnNewTitle(t *testing.T) {
	expDate := time.Now().AddDate(0, 0, 6).Format("2006-01-02")
	_, err := NewTitle(expDate, 100, 231654)
	if err != nil {
		t.Fail()
	}
}

func TestShouldCreateNewTitleWithEqualExpireDateAndCreateDate(t *testing.T) {
	expDate := time.Now().Format("2006-01-02")
	_, err := NewTitle(expDate, 100, 231654)
	if err != nil {
		t.Fail()
	}
}

func TestShouldFailWithCreateDateBiggerThanExpireDate(t *testing.T) {
	expDate := time.Now().AddDate(0, 0, -6).Format("2006-01-02")
	_, err := NewTitle(expDate, 100, 231654)
	if err == nil {
		t.Fail()
	}
}

func TestShouldFailWithAmountInCentsMinorThanOne(t *testing.T) {
	expDate := time.Now().AddDate(0, 0, 6).Format("2006-01-02")
	_, err := NewTitle(expDate, 0, 231654)
	if err == nil {
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
	e.Append(ErrorResponse{Code: "100", Message: "Hue"})
	test.ExpectTrue(len(e) == 2, t)
}

func TestShouldCreateNewSingleErrorCollection(t *testing.T) {
	e := NewSingleErrorCollection("200", "Hue2")
	test.ExpectTrue(len(e) == 1, t)
}
