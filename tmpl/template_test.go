package tmpl

import (
	"fmt"
	"testing"

	"time"

	"bitbucket.org/mundipagg/boletoapi/test"
)

func TestShouldTransformFromOriginToDestiny(t *testing.T) {
	tmp := `Ola {{.Username}} {{replace (today | brdate) "/" "."}}`
	b := New()
	final, err := b.From(struct{ Username string }{"username"}).To(tmp).Transform()
	test.ExpectNoError(err, t)
	date := time.Now().Format("02.01.2006")
	s := fmt.Sprintf("Ola username %s", date)
	test.ExpectTrue(final == s, t)
}

func TestShouldPadLeft(t *testing.T) {
	s := padLeft("5", "0", 5)
	test.ExpectTrue(s == "00005", t)
}

func TestShouldReturnString(t *testing.T) {
	test.ExpectTrue(toString(5) == "5", t)
}
func TestFormatDigitableLine(t *testing.T) {
	s := "34191123456789010111213141516171812345678901112"
	test.ExpectTrue(fmtDigitableLine(s) == "34191.12345 67890.101112 13141.516171 8 12345678901112", t)
}

func TestFormatCNPJ(t *testing.T) {
	s := "01000000000100"
	test.ExpectTrue(fmtCNPJ(s) == "01.000.000/0001-00", t)
}

func TestFormatCPF(t *testing.T) {
	s := "12312100100"
	test.ExpectTrue(fmtCPF(s) == "123.121.001-00", t)
}
