package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"

	controller "github.com/pierresantana/golang-restapi/controllers"
	. "github.com/pierresantana/golang-restapi/dao"
)

var dbURI = flag.String("dbURI", "", "URI for MongoDB instance")
var dbName = flag.String("dbName", "", "MongoDB database name")

var dao = ProductsDAO{}

func init() {
	fmt.Println("Golang Rest API")
	flag.Parse()

	if *dbURI == "" || *dbName == "" {
		flag.PrintDefaults()
		os.Exit(0)
	}

	dao.Server = *dbURI
	dao.Database = *dbName
	dao.Connect()
}

func main() {
	r := controller.Route()

	var port = ":3000"
	fmt.Println("Server running in port:", port)
	log.Fatal(http.ListenAndServe(port, r))
}
