package controller

import (
	"github.com/gorilla/mux"
)

func Route() *mux.Router {
	r := mux.NewRouter()

	r.HandleFunc("/products", ProductGetAll).Methods("GET")
	r.HandleFunc("/products/{id}", ProductGetByID).Methods("GET")
	r.HandleFunc("/products", ProductCreate).Methods("POST")
	r.HandleFunc("/products/{id}", ProductUpdate).Methods("PUT")
	r.HandleFunc("/products/{id}", ProductDelete).Methods("DELETE")

	return r
}
