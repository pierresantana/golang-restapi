package models

import "gopkg.in/mgo.v2/bson"

type Product struct {
	ID          bson.ObjectId `bson:"_id" json:"id"`
	Name        string        `bson:"name" json:"name"`
	Description string        `bson:"description" json:"description"`
	Price       float32       `bson:"price" json:"price"`
}
