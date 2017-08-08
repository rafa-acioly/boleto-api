package citibank

import (
	"github.com/mundipagg/boleto-api/validations"
	"github.com/mundipagg/boleto-api/models"
	"fmt"
	"errors"
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

func citiValidateAccountDigit(b interface{}) error {
	switch t := b.(type) {
	case *models.BoletoRequest:
		if len(t.Agreement.AccountDigit) < 1 || len(t.Agreement.AccountDigit) > 2{
			return errors.New(fmt.Sprintf("O digito da conta precisa ser preenchido."))
		}
		return nil
	default:
		return validations.InvalidType(t)
	}
}
