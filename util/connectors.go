package util

import (
	"bitbucket.org/mundipagg/boletoapi/log"

	"github.com/PMoneda/gonnie"
)

// SeqLogConector é um connector gonnie para logar no Seq
func SeqLogConector(next func(), e *gonnie.ExchangeMessage, out gonnie.Message, u gonnie.Uri, params ...interface{}) error {

	b := e.GetBody().(string)
	if b == "" {
		b = "Nenhum retorno do serviço"
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
