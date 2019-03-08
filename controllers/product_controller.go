package controller

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	. "github.com/pierresantana/golang-restapi/dao"
	. "github.com/pierresantana/golang-restapi/models"
	"gopkg.in/mgo.v2/bson"
)

var dao = ProductsDAO{}

func respondWithError(w http.ResponseWriter, code int, msg string) {
	respondWithJson(w, code, map[string]string{"error": msg})
}

func respondWithJson(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

func ProductGetAll(w http.ResponseWriter, r *http.Request) {
	products, err := dao.GetAll()
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJson(w, http.StatusOK, products)
}

func ProductGetByID(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	product, err := dao.GetByID(params["id"])
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid Product ID")
		return
	}
	respondWithJson(w, http.StatusOK, product)
}

func ProductCreate(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var product Product
	if err := json.NewDecoder(r.Body).Decode(&product); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	product.ID = bson.NewObjectId()
	if err := dao.Create(product); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJson(w, http.StatusCreated, product)
}

func ProductUpdate(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	params := mux.Vars(r)
	var product Product
	if err := json.NewDecoder(r.Body).Decode(&product); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	if err := dao.Update(params["id"], product); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJson(w, http.StatusOK, map[string]string{"result": product.Name + " atualizado com sucesso!"})
}

func ProductDelete(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	params := mux.Vars(r)
	if err := dao.Delete(params["id"]); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJson(w, http.StatusOK, map[string]string{"result": "success"})
}
