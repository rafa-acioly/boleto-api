package models

import (
	"fmt"
	"regexp"

	"bitbucket.org/mundipagg/boletoapi/util"
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
func (a *Agreement) IsAgencyValid() error {
	re := regexp.MustCompile("(\\D+)")
	ag := util.PadLeft(re.ReplaceAllString(a.Agency, ""), "0", 4)
	if len(ag) < 5 {
		a.Agency = ag
		return nil
	}
	return NewErrorResponse("MPAgency", "Agência inválida, deve conter até 4 dígitos")
}

// CalculateAgencyDigit calcula dígito da agência
func (a *Agreement) CalculateAgencyDigit(digitCalculator func(agency string) string) {
	re := regexp.MustCompile("(\\D+)")
	ad := re.ReplaceAllString(a.AgencyDigit, "")
	l := len(ad)
	if l < 2 && l > 0 {
		a.AgencyDigit = ad
	} else {
		a.AgencyDigit = digitCalculator(a.Agency)
	}
}

// IsAccountValid retorna se é uma conta válida
func (a *Agreement) IsAccountValid(accountLength int) error {
	re := regexp.MustCompile("(\\D+)")
	ac := util.PadLeft(re.ReplaceAllString(a.Account, ""), "0", uint(accountLength))
	if len(ac) < accountLength+1 {
		a.Account = ac
		return nil
	}
	return NewErrorResponse("MPAccount", fmt.Sprintf("Conta inválida, deve conter até %d dígitos", accountLength))
}

// IsAccountDigitValid retorna se o dígito da conta é válido
func (a Agreement) IsAccountDigitValid() (string, error) {
	re := regexp.MustCompile("(\\D+)")
	ad := re.ReplaceAllString(a.AccountDigit, "")
	l := len(ad)
	if l < 2 && l > 0 {
		return ad, nil
	}
	return "", NewErrorResponse("MPAccountDigit", "Dígito da conta inválido. Deve conter apenas um dígito.")
}

//CalculateAccountDigit calcula dígito da conta
func (a *Agreement) CalculateAccountDigit(digitCalculator func(agency, account string) string) {
	re := regexp.MustCompile("(\\D+)")
	ad := re.ReplaceAllString(a.AccountDigit, "")
	l := len(ad)
	if l < 2 && l > 0 {
		a.AccountDigit = ad
	}
	a.AccountDigit = digitCalculator(a.Agency, a.Account)
}
