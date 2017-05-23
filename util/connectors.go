package util

import (
	"fmt"
	"os"

	"bitbucket.org/mundipagg/boletoapi/log"

	"github.com/PMoneda/gonnie"
)

// SeqLogConector é um connector gonnie para logar no Seq
func SeqLogConector(next func(), e *gonnie.ExchangeMessage, out gonnie.Message, u gonnie.Uri, params ...interface{}) error {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Fudeu!")
			os.Exit(-20)
		}
	}()
	var b string
	switch t := e.GetBody().(type) {
	case string:
		if t == "" {
			b = "Nenhum retorno do serviço"
		} else {
			b = t
		}
	default:
		b = "Deu ruim"
	}
	if len(params) > 0 {
		l := params[0].(*log.Log)
		if u.GetOption("type") == "request" {
			l.Request(b, u.GetOption("url"), e.GetHeaderMap())
		}
		if u.GetOption("type") == "response" {
			l.Response(b, u.GetOption("url"))
		}
	}
	out <- e
	next()
	return nil
}
