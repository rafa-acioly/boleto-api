package models

import (
	"strings"
)

type Document struct {
	Type   DocumentType
	Number DocumentNumber
}

type DocumentType string

func (d DocumentType) IsCpf() bool {
	return strings.ToUpper(string(d)) == "CPF"
}

func (d DocumentType) IsCnpj() bool {
	return strings.ToUpper(string(d)) == "CNPJ"
}

type DocumentNumber string

func (d DocumentNumber) IsCpf() bool {
	return len(d) == 11
}

func (d DocumentNumber) IsCnpj() bool {
	return len(d) == 14
}
