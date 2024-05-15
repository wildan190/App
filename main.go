// Package main provides a simple REST API for managing products
package main

import (
	"encoding/json"
	"encoding/xml"
	"io/ioutil"
	"log"

	"monstercode/app"
	"net/http"
)

// main function as entry point
func main() {
	http.HandleFunc("/api/v1/products.json", productsJSONHandler)        // handle GET products in JSON format
	http.HandleFunc("/api/v1/products.xml", productsXMLHandler)          // handle GET products in XML format
	http.HandleFunc("/api/v1/add-products.json", addProductsJSONHandler) // handle POST products in JSON format
	http.HandleFunc("/api/v1/add-products.xml", addProductsXMLHandler)   // handle POST products in XML format
	log.Println("Starting server on :8081")
	if err := http.ListenAndServe(":8081", nil); err != nil {
		log.Fatalf("ListenAndServe: %v", err)
	}
}

// productsJSONHandler handles GET /api/v1/products.json
//
// Handle GET products in JSON format
func productsJSONHandler(w http.ResponseWriter, r *http.Request) {
	products := []app.Product{
		createProduct(1, "Product 1", 50000, 2022),
		createProduct(2, "Product 2", 75000, 2023),
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(products)
}

// productsXMLHandler handles GET /api/v1/products.xml
//
// Handle GET products in XML format
func productsXMLHandler(w http.ResponseWriter, r *http.Request) {
	products := []app.Product{
		createProduct(1, "Product 1", 50000, 2022),
		createProduct(2, "Product 2", 75000, 2023),
	}

	w.Header().Set("Content-Type", "text/xml")
	xml.NewEncoder(w).Encode(products)
}

// addProductsJSONHandler handles POST /api/v1/add-products.json
//
// Handle POST products in JSON format
func addProductsJSONHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var product app.Product
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Error reading request", http.StatusBadRequest)
		return
	}

	if err := json.Unmarshal(body, &product); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(product)
}

// addProductsXMLHandler handles POST /api/v1/add-products.xml
//
// Handle POST products in XML format
func addProductsXMLHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var product app.Product
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Error reading request", http.StatusBadRequest)
		return
	}

	if err := xml.Unmarshal(body, &product); err != nil {
		http.Error(w, "Invalid XML", http.StatusBadRequest)
		return
	}

	// Response in JSON format despite the input being XML
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(product)
}

// createProduct is a helper function to create products
//
// Helper function to create products
func createProduct(id int, name string, price int64, year int) app.Product {
	prod := app.Product{
		Price:    price,
		Category: &app.Category{},
	}
	prod.SetIDAndName(id, name)
	prod.SetYear(year)
	return prod
}
