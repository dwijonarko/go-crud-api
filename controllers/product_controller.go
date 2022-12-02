package controllers

import (
	"encoding/json"
	"go-crud-api/database"
	"go-crud-api/entities"
	"net/http"

	"github.com/gorilla/mux"
)

func GetProducts(w http.ResponseWriter, r *http.Request) {
	var products []entities.Product
	database.Instance.Find(&products)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(products)
}

func GetProductById(w http.ResponseWriter, r *http.Request) {
	productID := mux.Vars(r)["id"]
	var product entities.Product
	w.Header().Set("Content-Type", "application/json")
	database.Instance.First(&product, productID)
	if product.ID == 0 {
		json.NewEncoder(w).Encode("Product not found")
	} else {
		json.NewEncoder(w).Encode(product)
	}

}

func CreateProduct(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var product entities.Product
	json.NewDecoder(r.Body).Decode(&product)
	database.Instance.Create(&product)
	json.NewEncoder(w).Encode(product)
}

func UpdateProduct(w http.ResponseWriter, r *http.Request) {
	productID := mux.Vars(r)["id"]
	var product entities.Product
	w.Header().Set("Content-Type", "application/json")
	database.Instance.First(&product, productID)
	if product.ID == 0 {
		json.NewEncoder(w).Encode("Product not found")
	} else {
		json.NewDecoder(r.Body).Decode(&product)
		database.Instance.Save(&product)
		json.NewEncoder(w).Encode(product)
	}

}

func DeleteProduct(w http.ResponseWriter, r *http.Request) {
	productID := mux.Vars(r)["id"]
	var product entities.Product
	w.Header().Set("Content-Type", "application/json")
	database.Instance.First(&product, productID)
	if product.ID == 0 {
		json.NewEncoder(w).Encode("Product not found")
	} else {
		json.NewDecoder(r.Body).Decode(&product)
		database.Instance.Delete(&product, productID)
		json.NewEncoder(w).Encode("Product deleted")
	}
}
