package db

import "bitbucket.org/mundipagg/boletoapi/models"
import "bitbucket.org/mundipagg/boletoapi/cache"

type mock struct{}

//SaveBoleto salva o boleto num cache local em memoria
func (m *mock) SaveBoleto(id string, boleto models.BoletoView) error {
	cache.Set(id, boleto)
	return nil
}

//GetBoletoById retorna o boleto por id do cache em memoria
func (m *mock) GetBoletoByID(id string) (models.BoletoView, error) {
	return cache.Get(id).(models.BoletoView), nil
}
