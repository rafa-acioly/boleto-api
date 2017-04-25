package models

import (
	"regexp"
)

// Agreement afiliação do cliente com o bano
type Agreement struct {
	AgreementNumber int
	Wallet          int16
	WalletVariation int16
	Agency          string
	AgencyDigit     string
	Account         string
	AccountDigit    string
}

// IsAgencyValid retorna se é uma agência válida
func (a *Agreement) IsAgencyValid() bool {
	re := regexp.MustCompile("")
	re.ReplaceAllString(a.Agency, "")
	return len(a.Account) > 4
}

// IsAgencyDigitValid retorna se o dígito da agência é válido
func (a *Agreement) IsAgencyDigitValid() bool {
	return len(a.AgencyDigit) > 1
}
