package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	productController "github.com/pierresantana/golang-restapi/controllers/product"
	productRepository "github.com/pierresantana/golang-restapi/repositories/product"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var dbURI = flag.String("db-uri", "", "URI for MongoDB instance")
var dbName = flag.String("db-name", "", "MongoDB database name")

func main() {
	fmt.Println("Golang Rest API")
	flag.Parse()

	if *dbURI == "" || *dbName == "" {
		flag.PrintDefaults()
		os.Exit(0)
	}

	session, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(*dbURI))
	if err != nil {
		log.Fatalf("fail to connect to %s: %s", *dbURI, err)
	}
	db := session.Database((*dbName))
	r := mux.NewRouter()
	productController.Register(r, productRepository.NewMongoProductRepository(db))

	var port = ":3000"
	fmt.Println("Server running in port:", port)
	log.Fatal(http.ListenAndServe(port, r))
}
