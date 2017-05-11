package db

import (
	"sync"

	"bitbucket.org/mundipagg/boletoapi/models"
)

//DB é a interface basica para os métodos de persistência
type DB interface {
	SaveBoleto(models.BoletoView) error
	GetBoletoByID(string) (models.BoletoView, error)
	Close()
}

var db DB = nil
var _create sync.Mutex
var err error

//GetDB retorna o objeto concreto que implementa as funções de persistência
func GetDB() (DB, error) {
	return CreateMongo()
}
