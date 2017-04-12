package models

import (
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

// IsCpf verifica se é um Cpf válido
func (d DocumentNumber) IsCpf() bool {
	return len(d) == 11
}

// IsCnpj verifica se é um Cnpj válido
func (d DocumentNumber) IsCnpj() bool {
	return len(d) == 14
}
