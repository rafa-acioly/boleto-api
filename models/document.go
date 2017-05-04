package models

import (
	"regexp"
	"strings"
)

// Document nó com o tipo de documento e número do documento
type Document struct {
	Type   DocumentType
	Number DocumentNumber
}

// DocumentType o tipo de documento pode ser CPF ou CNPJ
type DocumentType string

// IsCpf diz se o DocumentType é um CPF
func (d DocumentType) IsCpf() bool {
	return strings.ToUpper(string(d)) == "CPF"
}

// IsCnpj diz se o DocumentType é um CNPJ
func (d DocumentType) IsCnpj() bool {
	return strings.ToUpper(string(d)) == "CNPJ"
}

// DocumentNumber o número do documento, poder ser um CPF ou CNPJ
type DocumentNumber string

// ValidateCPF verifica se é um CPF válido
func (d *Document) ValidateCPF() error {
	re := regexp.MustCompile("(\\D+)")
	cpf := re.ReplaceAllString(string(d.Number), "")
	if len(cpf) == 11 {
		d.Number = DocumentNumber(cpf)
		return nil
	}
	return NewErrorResponse("MPDocumentNumber", "CPF inválido")
}

// ValidateCNPJ verifica se é um CNPJ válido
func (d *Document) ValidateCNPJ() error {
	re := regexp.MustCompile("(\\D+)")
	cnpj := re.ReplaceAllString(string(d.Number), "")
	if len(cnpj) == 14 {
		d.Number = DocumentNumber(cnpj)
		return nil
	}
	return NewErrorResponse("MPDocumentNumber", "CNPJ inválido")
}
