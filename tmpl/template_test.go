package tmpl

import (
	"fmt"
	"testing"

	"bitbucket.org/mundipagg/boletoapi/test"
)

func TestShouldTransformFromOriginToDestiny(t *testing.T) {
	tmp := `Ola {{.Username}} {{replace (today | brdate) "/" "."}}`
	b := New()
	final, err := b.From(struct{ Username string }{"username"}).To(tmp).Transform()
	fmt.Println(final)
	test.ExpectNoError(err, t)
	test.ExpectTrue(final == "Ola username 13.04.2017", t)

}
