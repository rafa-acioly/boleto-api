package db

import (
	"sync"

	"bitbucket.org/mundipagg/boletoapi/config"
	"bitbucket.org/mundipagg/boletoapi/models"
)

//DB é a interface basica para os métodos de persistência
type DB interface {
	SaveBoleto(models.BoletoView) error
	GetBoletoByID(string) (models.BoletoView, error)
	Close()
}

var db DB
var _createOnce sync.Once

//GetDB retorna o objeto concreto que implementa as funções de persistência
func GetDB() DB {
	_createOnce.Do(func() {
		if config.Get().MockMode {
			db = new(mock)
		}
		db = CreateMongo()
	})
	return db
}
