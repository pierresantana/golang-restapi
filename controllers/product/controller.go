package product

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/pierresantana/golang-restapi/models"
	productRepository "github.com/pierresantana/golang-restapi/repositories/product"
	"github.com/pierresantana/golang-restapi/services/product"
	"github.com/pierresantana/golang-restapi/utils"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var service *product.ProductsService

func Register(r *mux.Router, repository productRepository.ProductRepository) *mux.Router {
	service = product.NewService(repository)

	r.HandleFunc("/products", ProductGetAll).Methods("GET")
	r.HandleFunc("/products/{id}", ProductGetByID).Methods("GET")
	r.HandleFunc("/products", ProductCreate).Methods("POST")
	r.HandleFunc("/products/{id}", ProductUpdate).Methods("PUT")
	r.HandleFunc("/products/{id}", ProductDelete).Methods("DELETE")

	return r
}

func ProductGetAll(w http.ResponseWriter, r *http.Request) {
	products, err := service.GetAll(r.Context())
	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	if products == nil {
		products = []*models.Product{}
	}
	utils.RespondWithJson(w, http.StatusOK, products)
}

func ProductGetByID(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	product, err := service.GetByID(r.Context(), params["id"])
	if err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, "Invalid Product ID")
		return
	}
	utils.RespondWithJson(w, http.StatusOK, product)
}

func ProductCreate(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var product models.Product
	if err := json.NewDecoder(r.Body).Decode(&product); err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	product.ID = primitive.NewObjectID()
	if err := service.Create(r.Context(), &product); err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	utils.RespondWithJson(w, http.StatusCreated, product)
}

func ProductUpdate(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	params := mux.Vars(r)
	var product models.Product
	var err error
	if err := json.NewDecoder(r.Body).Decode(&product); err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	product.ID, err = primitive.ObjectIDFromHex(params["id"])
	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	if err := service.Update(r.Context(), params["id"], &product); err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	utils.RespondWithJson(w, http.StatusOK, product)
}

func ProductDelete(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	params := mux.Vars(r)
	if err := service.Delete(r.Context(), params["id"]); err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
