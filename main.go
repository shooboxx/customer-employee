package main

import (
	"log"
	"net/http"

	"practice/app"

	"github.com/gorilla/mux"
)

// Main function
func main() {
	// Init router
	r := mux.NewRouter()
	h := app.NewHandler()

	// Route handles & endpoints
	r.HandleFunc("/customers", h.GetCustomers).Methods("GET")
	r.HandleFunc("/products", h.GetProducts).Methods("GET")
	r.HandleFunc("/customer/{id}", h.UpdateCustomer).Methods("POST")
	r.HandleFunc("/product/{id}", h.UpdateProduct).Methods("POST")

	// Start server
	log.Fatal(http.ListenAndServe(":8080", r))
}
