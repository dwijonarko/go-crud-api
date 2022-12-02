package main

import (
	"fmt"
	"go-crud-api/controllers"
	"go-crud-api/database"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	LoadAppConfig() //load configuration

	// initialize database
	database.Connect(AppConfig.ConnectionString)
	database.Migrate()

	//initialize routing
	router := mux.NewRouter().StrictSlash(true)

	//register router
	ProductHandlers(router)

	//Start server
	log.Println(fmt.Sprintf("Starting server on port %s", AppConfig.Port))
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%v", AppConfig.Port), router))
}

func ProductHandlers(router *mux.Router) {
	router.HandleFunc("/api/products", controllers.GetProducts).Methods(http.MethodGet)
	router.HandleFunc("/api/products/{id}", controllers.GetProductById).Methods(http.MethodGet)
	router.HandleFunc("/api/products", controllers.CreateProduct).Methods(http.MethodPost)
	router.HandleFunc("/api/products/{id}", controllers.UpdateProduct).Methods(http.MethodPut)
	router.HandleFunc("/api/products/{id}", controllers.DeleteProduct).Methods(http.MethodDelete)
}
