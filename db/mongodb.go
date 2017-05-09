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
	session *mgo.Session
	m       sync.RWMutex
}

//CreateMongo cria uma nova intancia de conexão com o mongodb
func CreateMongo() (DB, error) {
	db := new(mongoDb)
	var err error
	db.session, err = mgo.Dial(config.Get().MongoURL)
	if err != nil {
		return nil, models.NewInternalServerError(err.Error(), "Falha ao conectar com o banco de dados")
	}
	return db, nil
}

//SaveBoleto salva um boleto no mongoDB
func (e *mongoDb) SaveBoleto(boleto models.BoletoView) error {
	var err error
	e.m.Lock()
	defer e.m.Unlock()
	if e.session == nil {
		return models.NewInternalServerError("Database error", "Falha ao conectar com o banco de dados")
	}
	c := e.session.DB("boletoapi").C("boletos")
	err = c.Insert(boleto)
	return err
}

//GetBoletoById busca um boleto pelo ID que vem na URL
func (e *mongoDb) GetBoletoByID(id string) (models.BoletoView, error) {
	c := e.session.DB("boletoapi").C("boletos")
	result := models.BoletoView{}
	err := c.Find(bson.M{"id": id}).One(&result)
	if err != nil {
		return models.BoletoView{}, err
	}
	return result, nil
}

func (e *mongoDb) Close() {
	fmt.Println("Close Database Connection")
	e.session.Close()
}
