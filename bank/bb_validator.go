package bank

import (
	"strconv"

	"bitbucket.org/mundipagg/boletoapi/models"
)

var bb bankBB

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

func bbAgencyDigitCalculator(agency string) string {
	multiplier := []int{5, 4, 3, 2}
	return modElevenCalculator(agency, multiplier)
}

func bbAccountDigitCalculator(agency, account string) string {
	multiplier := []int{9, 8, 7, 6, 5, 4, 3, 2}
	return modElevenCalculator(account, multiplier)
}

func bbValidateAgencyAndDigit(b interface{}) error {
	switch t := b.(type) {
	case *models.BoletoRequest:
		err := t.Agreement.IsAgencyValid()
		if err != nil {
			return err
		}
		t.Agreement.CalculateAgencyDigit(bbAgencyDigitCalculator)
		return nil
	default:
		return invalidType(t)
	}
}

func invalidType(t interface{}) error {
	//tp := reflect.TypeOf(t)
	//fmt.Println("O tipo vindo do request não é models.BoletoRequest mas sim: " + tp.String())
	return models.NewErrorResponse("MP500", "Tipo inválido")
}

func bbValidateAccountAndDigit(b interface{}) error {
	switch t := b.(type) {
	case *models.BoletoRequest:
		err := t.Agreement.IsAccountValid(8)
		if err != nil {
			return err
		}
		t.Agreement.CalculateAccountDigit(bbAccountDigitCalculator)
		return nil
	default:
		return invalidType(t)
	}
}

func bbValidateOurNumber(b interface{}) error {
	switch t := b.(type) {
	case *models.BoletoRequest:
		if t.Title.OurNumber > 9999999999 {
			return models.NewErrorResponse("MPOurNumber", "Nosso número inválido")
		}
		return nil
	default:
		return invalidType(t)
	}
}
