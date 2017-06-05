package util

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestEncryptDecrypt(t *testing.T) {
	Convey("Deve encriptar o texto", t, func() {
		a := Encrypt("asd")
		So(a, ShouldNotEqual, "asd")
		Convey("Deve desencriptar o texto", func() {
			x := Decrypt(a)
			So(x, ShouldEqual, "asd")
		})
	})
}

func TestBase64EncodeDecode(t *testing.T) {
	Convey("Deve encodar em Base64 o texto", t, func() {
		a := Base64("asd")
		So(a, ShouldNotEqual, "asd")
		Convey("Deve desencodar o texto Base64", func() {
			x := Base64Decode(a)
			So(x, ShouldEqual, "asd")
		})
	})
}
