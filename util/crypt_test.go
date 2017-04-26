package util

import (
	"testing"

	"bitbucket.org/mundipagg/boletoapi/test"
)

func TestEncryptDecrypt(t *testing.T) {
	test.ExpectTrue(Decrypt(Encrypt("asd")) == "asd", t)
}
