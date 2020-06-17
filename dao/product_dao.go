package dao

import (
	"log"

	. "github.com/pierresantana/golang-restapi/models"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type ProductsDAO struct {
	Server   string
	Database string
}

var db *mgo.Database

const (
	COLLECTION = "products"
)

func (m *ProductsDAO) Connect() {
	session, err := mgo.Dial(m.Server)
	if err != nil {
		log.Fatal(err)
	}
	db = session.DB(m.Database)
}

func (m *ProductsDAO) GetAll() ([]Product, error) {
	var product []Product
	err := db.C(COLLECTION).Find(bson.M{}).All(&product)
	return product, err
}

func (m *ProductsDAO) GetByID(id string) (Product, error) {
	var product Product
	err := db.C(COLLECTION).FindId(bson.ObjectIdHex(id)).One(&product)
	return product, err
}

func (m *ProductsDAO) Create(product Product) error {
	err := db.C(COLLECTION).Insert(&product)
	return err
}

func (m *ProductsDAO) Delete(id string) error {
	err := db.C(COLLECTION).RemoveId(bson.ObjectIdHex(id))
	return err
}

func (m *ProductsDAO) Update(id string, product Product) error {
	err := db.C(COLLECTION).UpdateId(bson.ObjectIdHex(id), &product)
	return err
}
