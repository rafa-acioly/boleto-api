package bank

import (
	"testing"

	"bitbucket.org/mundipagg/boletoapi/test"
)

func TestShouldCalculateAgencyDigitFromBb(t *testing.T) {
	test.ExpectTrue(agencyDigitCalculator("0137") == "6", t)

	test.ExpectTrue(agencyDigitCalculator("3418") == "5", t)

	test.ExpectTrue(agencyDigitCalculator("3324") == "3", t)

	test.ExpectTrue(agencyDigitCalculator("5797") == "5", t)
}
