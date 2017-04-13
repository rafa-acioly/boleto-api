package test

import "testing"

// ExpectNoError falha o teste se e != nil
func ExpectNoError(e error, t *testing.T) {
	if e != nil {
		t.Fail()
	}
}

// ExpectTrue falha o teste caso a condição não seja verdadeira
func ExpectTrue(condition bool, t *testing.T) {
	if !condition {
		t.Fail()
	}
}

// ExpectFalse falha o teste caso a condição não seja falsa
func ExpectFalse(condition bool, t *testing.T) {
	if condition {
		t.Fail()
	}
}

// ExpectNil falha o teste caso obj seja diferente de nil
func ExpectNil(obj interface{}, t *testing.T) {
	if obj != nil {
		t.Fail()
	}
}
