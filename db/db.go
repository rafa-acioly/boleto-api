package db

import (
	"bitbucket.org/mundipagg/boletoapi/config"
	"bitbucket.org/mundipagg/boletoapi/models"
)

//DB é a interface basica para os métodos de persistência
type DB interface {
	SaveBoleto(string, models.BoletoView) error
	GetBoletoByID(string) (models.BoletoView, error)
}

//GetDB retorna o objeto concreto que implementa as funções de persistência
func GetDB() DB {
	if config.Get().MockMode {
		return new(mock)
	}
	return new(elasticDb)
}
