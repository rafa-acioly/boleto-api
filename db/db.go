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

var db DB = nil
var _create sync.Mutex

//GetDB retorna o objeto concreto que implementa as funções de persistência
func GetDB() (DB, error) {
	var err error
	_create.Lock()
	defer _create.Unlock()
	if db != nil {
		return db, nil
	}
	if config.Get().MockMode {
		db = new(mock)
	} else {
		db, err = CreateMongo()
	}
	return db, err
}
