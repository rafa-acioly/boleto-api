package tmpl

import (
	"testing"

	"bitbucket.org/mundipagg/boletoapi/test"
)

func TestShouldTransformFromOriginToDestiny(t *testing.T) {
	tmp := "Ola {{.Username}}"
	b := New()
	final, err := b.From(struct{ Username string }{"username"}).To(tmp).Transform()
	test.ExpectNoError(err, t)
	test.ExpectTrue(final == "Ola username", t)

}
