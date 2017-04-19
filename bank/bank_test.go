package bank

import (
	"testing"

	"bitbucket.org/mundipagg/boletoapi/models"
	"bitbucket.org/mundipagg/boletoapi/test"
)

func TestShouldExecuteBBStrategy(t *testing.T) {
	var bb models.BankNumber = models.BancoDoBrasil
	bank, err := Get(bb)
	number := bank.GetBankNumber()
	test.ExpectNoError(err, t)
	test.ExpectTrue(number.IsBankNumberValid(), t)
	test.ExpectTrue(number == models.BancoDoBrasil, t)
}
