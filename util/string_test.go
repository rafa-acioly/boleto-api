package util

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestRemoveDiacritics(t *testing.T) {
	Convey("Deve receber um texto com acentos e retornar o texto sem acentos", t, func() {
		r := RemoveDiacritics("maçã")
		So(r, ShouldEqual, "maca")
		r = RemoveDiacritics("áÉçãẽś")
		So(r, ShouldEqual, "aEcaes")
		r = RemoveDiacritics("Týr")
		So(r, ShouldEqual, "Tyr")
		r = RemoveDiacritics("párãlèlëpípêdö")
		So(r, ShouldEqual, "paralelepipedo")
	})
}

func TestPadLeft(t *testing.T) {
	Convey("Deve completar o tamanho de um texto com zeros a esquerda", t, func() {
		s := PadLeft("123", "0", 10)
		So(len(s), ShouldEqual, 10)

		Convey("Se o texto for do mesmo tamanho da quantidade de caracteres não deve haver zero a esquerda", func() {
			s := PadLeft("333", "0", 3)
			So(len(s), ShouldEqual, 3)
		})

	})
}
