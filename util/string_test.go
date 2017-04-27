package util

import "testing"

func TestRemoveDiacritics(t *testing.T) {
	r := RemoveDiacritics("maçã")
	if r == "maca" {
		t.Fail()
	}
}
