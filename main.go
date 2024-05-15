package main

import (
	"encoding/json"
	"encoding/xml"
	"io/ioutil"
	"log"
	"monstercode/app"
	"net/http"
)

// handleProductsJSON adalah handler function untuk endpoint /api/v1/products.json
func handleProductsJSON(w http.ResponseWriter, r *http.Request) {
	// Definisikan slice dari []app.Product
	products := []app.Product{}

	// Buat 2 produk dan tambahkan ke slice
	product1 := app.Product{
		Price:    10000,
		Category: app.Category{},
	}
	product1.SetIDAndName(1, "Product 1")
	product1.SetYear(2022)
	product1.Category.SetIDAndName(1, "Category 1")

	product2 := app.Product{
		Price:    20000,
		Category: app.Category{},
	}
	product2.SetIDAndName(2, "Product 2")
	product2.SetYear(2023)
	product2.Category.SetIDAndName(2, "Category 2")

	products = append(products, product1, product2)

	// Set response header Content-Type ke application/json
	w.Header().Set("Content-Type", "application/json")
	// Encode slice products ke JSON dan tulis ke response writer
	json.NewEncoder(w).Encode(products)
}

// handleProductsXML adalah handler function untuk endpoint /api/v1/products.xml
func handleProductsXML(w http.ResponseWriter, r *http.Request) {
	// Definisikan slice dari []app.Product
	products := []app.Product{}

	// Buat 2 produk dan tambahkan ke slice
	product1 := app.Product{
		Price:    10000,
		Category: app.Category{},
	}
	product1.SetIDAndName(1, "Product 1")
	product1.SetYear(2022)
	product1.Category.SetIDAndName(1, "Category 1")

	product2 := app.Product{
		Price:    20000,
		Category: app.Category{},
	}
	product2.SetIDAndName(2, "Product 2")
	product2.SetYear(2023)
	product2.Category.SetIDAndName(2, "Category 2")

	products = append(products, product1, product2)

	// Set response header Content-Type ke text/xml
	w.Header().Set("Content-Type", "text/xml")
	// Encode slice products ke XML dan tulis ke response writer
	xml.NewEncoder(w).Encode(products)
}

// handleAddProductJSON adalah handler function untuk endpoint /api/v1/add-products.json
func handleAddProductJSON(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	// Baca request body
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Unable to read request body", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	// Unmarshal JSON ke struct Product
	var product app.Product
	err = json.Unmarshal(body, &product)
	if err != nil {
		http.Error(w, "Unable to parse JSON", http.StatusBadRequest)
		return
	}

	// Marshal kembali struct Product ke JSON untuk dijadikan response
	response, err := json.Marshal(product)
	if err != nil {
		http.Error(w, "Unable to marshal JSON", http.StatusInternalServerError)
		return
	}

	// Set response header Content-Type ke application/json
	w.Header().Set("Content-Type", "application/json")
	// Tulis response JSON
	w.Write(response)
}

// handleAddProductXML adalah handler function untuk endpoint /api/v1/add-products.xml
func handleAddProductXML(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	// Baca request body
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Unable to read request body", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	// Unmarshal XML ke struct Product
	var product app.Product
	err = xml.Unmarshal(body, &product)
	if err != nil {
		http.Error(w, "Unable to parse XML", http.StatusBadRequest)
		return
	}

	// Marshal kembali struct Product ke JSON untuk dijadikan response
	response, err := json.Marshal(product)
	if err != nil {
		http.Error(w, "Unable to marshal JSON", http.StatusInternalServerError)
		return
	}

	// Set response header Content-Type ke application/json
	w.Header().Set("Content-Type", "application/json")
	// Tulis response JSON
	w.Write(response)
}

func main() {
	// Daftarkan endpoint dan handler-nya
	http.HandleFunc("/api/v1/products.json", handleProductsJSON)
	http.HandleFunc("/api/v1/products.xml", handleProductsXML)
	http.HandleFunc("/api/v1/add-products.json", handleAddProductJSON)
	http.HandleFunc("/api/v1/add-products.xml", handleAddProductXML)

	// Listen dan serve HTTP server pada port 8081
	log.Println("Starting server on :8081")
	if err := http.ListenAndServe(":8081", nil); err != nil {
		log.Fatalf("ListenAndServe: %v", err)
	}
}
