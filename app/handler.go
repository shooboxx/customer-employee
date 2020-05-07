package app

import (
	"fmt"
	"net/http"
	// "github.com/gorilla/mux"
	// "encoding/json"
)

// Get customers
func (h *Handler) GetCustomers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	fmt.Println("getCustomers route")
	// json.NewEncoder(w).Encode()
}

// Get products
func (h *Handler) GetProducts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	fmt.Println("getProducts route")
	// json.NewEncoder(w).Encode()
}

// Update customer details
func (h *Handler) UpdateCustomer(w http.ResponseWriter, r *http.Request) {
	fmt.Println("updateCustomer route")
	w.Header().Set("Content-Type", "application/json")
	// params := mux.Vars(r)
}

// Update product details
func (h *Handler) UpdateProduct(w http.ResponseWriter, r *http.Request) {
	fmt.Println("updateProduct route")
	w.Header().Set("Content-Type", "application/json")
	// params := mux.Vars(r)
}

// Delete customer
func (h *Handler) DeleteCustomer(w http.ResponseWriter, r *http.Request) {
	fmt.Println("deleteCustomer route")
	w.Header().Set("Content-Type", "application/json")
	// params := mux.Vars(r)
	// json.NewEncoder(w).Encode()
}

// Delete product
func (h *Handler) DeleteProduct(w http.ResponseWriter, r *http.Request) {
	fmt.Println("deleteProduct route")
	w.Header().Set("Content-Type", "application/json")
	// params := mux.Vars(r)
	// json.NewEncoder(w).Encode()
}
