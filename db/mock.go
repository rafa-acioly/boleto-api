package db

import (
	"errors"

	"bitbucket.org/mundipagg/boletoapi/cache"
	"bitbucket.org/mundipagg/boletoapi/models"
)

type mock struct{}

//SaveBoleto salva o boleto num cache local em memoria
func (m *mock) SaveBoleto(boleto models.BoletoView) error {
	cache.Set(boleto.ID, boleto)
	return nil
}

//GetBoletoById retorna o boleto por id do cache em memoria
func (m *mock) GetBoletoByID(id string) (models.BoletoView, error) {
	c, ok := cache.Get(id)
	if !ok {
		return models.BoletoView{}, errors.New("Boleto não encontrado")
	}
	return c.(models.BoletoView), nil
}

func (m *mock) Close() {}
