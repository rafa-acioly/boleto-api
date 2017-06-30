package bank

import (
	"testing"

	"github.com/mundipagg/boleto-api/models"
	"github.com/mundipagg/boleto-api/test"
)

func TestShouldExecuteBBStrategy(t *testing.T) {
	var bb models.BankNumber = models.BancoDoBrasil
	bank, err := Get(bb)
	number := bank.GetBankNumber()
	test.ExpectNoError(err, t)
	test.ExpectTrue(number.IsBankNumberValid(), t)
	test.ExpectTrue(number == models.BancoDoBrasil, t)
}
