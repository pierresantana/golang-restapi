package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	. "github.com/pierresantana/golang-restapi/config"
	. "github.com/pierresantana/golang-restapi/dao"
	productrouter "github.com/pierresantana/golang-restapi/router"
)

var dao = ProductsDAO{}
var config = Config{}

func init() {
	config.Read()

	dao.Server = config.Server
	dao.Database = config.Database
	dao.Connect()
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/products", productrouter.GetAll).Methods("GET")
	r.HandleFunc("/products/{id}", productrouter.GetByID).Methods("GET")
	r.HandleFunc("/products", productrouter.Create).Methods("POST")
	r.HandleFunc("/products/{id}", productrouter.Update).Methods("PUT")
	r.HandleFunc("/products/{id}", productrouter.Delete).Methods("DELETE")

	var port = ":8080"
	fmt.Println("Server running in port:", port)
	log.Fatal(http.ListenAndServe(port, r))
}
