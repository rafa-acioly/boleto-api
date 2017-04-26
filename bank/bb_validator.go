package bank

import (
	"strconv"

	"bitbucket.org/mundipagg/boletoapi/models"
)

func agencyDigitCalculator(agency string) string {
	multiplier := [4]int{5, 4, 3, 2}
	sum := 0

	for idx, c := range agency {
		i, _ := strconv.Atoi(string(c))

		sum += i * multiplier[idx]
	}

	digit := 11 - sum%11

	if digit == 10 {
		return "X"
	}

	if digit == 11 {
		return "0"
	}

	return strconv.Itoa(digit)
}

func validateAgency(b *models.BoletoRequest) error {
	err := b.Agreement.IsAgencyValid()
	return err
}

func validateAgencyDigit(b *models.BoletoRequest) error {
	return nil
}

func hue(boleto *models.BoletoRequest) models.Errors {
	err := models.NewEmptyErrorCollection()
	// if models.IsAgencyValid(&boleto.Agreement) {
	// 	if !models.IsAgencyDigitValid(&boleto.Agreement) {
	// 		boleto.Agreement.AgencyDigit = agencyDigitCalculator(boleto.Agreement.Agency)
	// 	}
	// } else {
	// 	err.Append("MPBB001", "Agência inválida")
	// }

	account, e := boleto.Agreement.IsAccountValid(8)
	if e != nil {
		ex, _ := e.(models.ErrorInterface)
		err.Append(ex.ErrorCode(), ex.Error())
	} else {
		boleto.Agreement.Account = account
		if ad, ed := boleto.Agreement.IsAccountDigitValid(); ed == nil {
			boleto.Agreement.AccountDigit = ad
		} else {
			// TODO: Fazer lógica para calcular dígito da conta
			ex, _ := e.(models.ErrorInterface)
			err.Append(ex.ErrorCode(), ex.Error())
		}
	}
	return err
}
