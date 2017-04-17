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
	fmt.Println(final)
	test.ExpectNoError(err, t)
	date := time.Now().Format("02.01.2006")
	s := fmt.Sprintf("Ola username %s", date)
	test.ExpectTrue(final == s, t)
}

func TestShouldPadLeft(t *testing.T) {
	s := padLeft("5", "0", 5)
	test.ExpectTrue(s == "00005", t)
}
