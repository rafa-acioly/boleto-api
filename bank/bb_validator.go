package bank

import (
	"strconv"

	"bitbucket.org/mundipagg/boletoapi/models"
)

func modElevenCalculator(a string, m []int) string {
	sum := 0

	for idx, c := range a {
		i, _ := strconv.Atoi(string(c))

		sum += i * m[idx]
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

func agencyDigitCalculator(agency string) string {
	multiplier := []int{5, 4, 3, 2}
	return modElevenCalculator(agency, multiplier)
}

func accountDigitCalculator(agency, account string) string {
	multiplier := []int{9, 8, 7, 6, 5, 4, 3, 2}
	return modElevenCalculator(account, multiplier)
}

func validateAgencyAndDigit(b *models.BoletoRequest) error {
	err := b.Agreement.IsAgencyValid()
	if err != nil {
		return err
	}
	b.Agreement.CalculateAgencyDigit(agencyDigitCalculator)
	return nil
}

func validateAccountAndDigit(b *models.BoletoRequest) error {
	err := b.Agreement.IsAccountValid(8)
	if err != nil {
		return err
	}
	b.Agreement.CalculateAccountDigit(accountDigitCalculator)
	return nil
}
