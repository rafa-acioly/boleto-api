package util

import (
	"fmt"

	"github.com/PMoneda/gonnie"
)

func SeqLogConector(next func(), e *gonnie.ExchangeMessage, out gonnie.Message, u gonnie.Uri, params ...interface{}) error {
	fmt.Println("E para logar no SEQ")
	out <- e
	next()
	return nil
}
