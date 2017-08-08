package citibank

import (
	"github.com/mundipagg/boleto-api/validations"
	"github.com/mundipagg/boleto-api/models"
)

func citiValidateAgency(b interface{}) error {
	switch t := b.(type) {
	case *models.BoletoRequest:
		err := t.Agreement.IsAgencyValid()
		if err != nil {
			return err
		}
		return nil
	default:
		return validations.InvalidType(t)
	}
}

func citiValidateAccount(b interface{}) error {
	switch t := b.(type) {
	case *models.BoletoRequest:
		err := t.Agreement.IsAccountValid(9)
		if err != nil {
			return err
		}
		return nil
	default:
		return validations.InvalidType(t)
	}
}

func citiValidateDigitQuantity(b interface{}) error {
	switch t := b.(type) {
	case *models.BoletoRequest:
		err := t.Agreement.IsQuantityDigitValidade(1)
		if err != nil {
			return err
		}
		return nil
	default:
		return validations.InvalidType(t)
	}
}
