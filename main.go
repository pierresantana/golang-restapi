package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	controller "github.com/pierresantana/golang-restapi/controllers"
	. "github.com/pierresantana/golang-restapi/dao"
)

const (
	dbUrl  = "DB_URL"
	dbName = "DB_NAME"
)

var dao = ProductsDAO{}

func init() {
	config := dbConfig()
	dao.Server = config[dbUrl]
	dao.Database = config[dbName]
	dao.Connect()
}

func dbConfig() map[string]string {
	conf := make(map[string]string)
	url, ok := os.LookupEnv(dbUrl)
	if !ok {
		panic("DB_URL environment variable required but not set")
	}
	name, ok := os.LookupEnv(dbName)
	if !ok {
		panic("DB_NAME environment variable required but not set")
	}
	conf[dbUrl] = url
	conf[dbName] = name
	return conf
}

func main() {
	r := controller.Route()

	var port = ":3000"
	fmt.Println("Server running in port:", port)
	log.Fatal(http.ListenAndServe(port, r))
}
