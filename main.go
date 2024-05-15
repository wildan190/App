package main

import (
	"encoding/json"
	"encoding/xml"
	"io/ioutil"
	"log"
	"monstercode/app"
	"net/http"
)

func main() {
	http.HandleFunc("/api/v1/products.json", productsJSONHandler)        // JSON endpoint for getting products
	http.HandleFunc("/api/v1/products.xml", productsXMLHandler)          // XML endpoint for getting products
	http.HandleFunc("/api/v1/add-products.json", addProductsJSONHandler) // JSON endpoint for adding products
	http.HandleFunc("/api/v1/add-products.xml", addProductsXMLHandler)   // XML endpoint for adding products
	log.Println("Starting server on :8081")
	if err := http.ListenAndServe(":8081", nil); err != nil {
		log.Fatalf("ListenAndServe: %v", err)
	}
}

func productsJSONHandler(w http.ResponseWriter, r *http.Request) {
	products := []app.Product{
		createProduct(1, "Product 1", 50000, 2022),
		createProduct(2, "Product 2", 75000, 2023),
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(products)
}

func productsXMLHandler(w http.ResponseWriter, r *http.Request) {
	products := []app.Product{
		createProduct(1, "Product 1", 50000, 2022),
		createProduct(2, "Product 2", 75000, 2023),
	}

	w.Header().Set("Content-Type", "text/xml")
	xml.NewEncoder(w).Encode(products)
}

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
