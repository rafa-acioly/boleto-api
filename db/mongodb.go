package db

import (
	"fmt"
	"sync"

	"bitbucket.org/mundipagg/boletoapi/config"
	"bitbucket.org/mundipagg/boletoapi/models"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type mongoDb struct {
	m sync.RWMutex
}

var dbName = "boletoapi"

//CreateMongo cria uma nova intancia de conexão com o mongodb
func CreateMongo() (DB, error) {
	db := new(mongoDb)
	if config.Get().MockMode {
		dbName = "boletoapi_mock"
	}
	return db, nil
}

//SaveBoleto salva um boleto no mongoDB
func (e *mongoDb) SaveBoleto(boleto models.BoletoView) error {
	var err error
	e.m.Lock()
	defer e.m.Unlock()
	session, err := mgo.Dial(config.Get().MongoURL)
	if err != nil {
		return models.NewInternalServerError(err.Error(), "Falha ao conectar com o banco de dados")
	}
	defer session.Close()
	c := session.DB(dbName).C("boletos")
	err = c.Insert(boleto)
	return err
}

//GetBoletoById busca um boleto pelo ID que vem na URL
func (e *mongoDb) GetBoletoByID(id string) (models.BoletoView, error) {
	e.m.Lock()
	defer e.m.Unlock()
	result := models.BoletoView{}
	session, err := mgo.Dial(config.Get().MongoURL)
	if err != nil {
		return result, models.NewInternalServerError(err.Error(), "Falha ao conectar com o banco de dados")
	}
	defer session.Close()
	c := session.DB(dbName).C("boletos")
	errF := c.Find(bson.M{"id": id}).One(&result)
	if errF != nil {
		return models.BoletoView{}, err
	}
	return result, nil
}

func (e *mongoDb) Close() {
	fmt.Println("Close Database Connection")
}
