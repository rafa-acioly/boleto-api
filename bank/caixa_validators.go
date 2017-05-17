package bank

import (
	"bitbucket.org/mundipagg/boletoapi/models"
)

func caixaAccountDigitCalculator(agency, account string) string {
	multiplier := []int{9, 8, 7, 6, 5, 4, 3, 2}
	return modElevenCalculator(account, multiplier)
}

func caixaValidateAccountAndDigit(b interface{}) error {
	switch t := b.(type) {
	case *models.BoletoRequest:
		//TODO Validar a regra certinho da Caixa
		err := t.Agreement.IsAccountValid(11)
		if err != nil {
			return err
		}
		return nil
	default:
		return invalidType(t)
	}
}
