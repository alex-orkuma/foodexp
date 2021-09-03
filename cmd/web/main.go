package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
)

func dasboard(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Welcome to my dashboard..."))
}

// Add products handler
func addProduct(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodPost {

		w.Header().Set("Allow", http.MethodPost)
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}
	w.Write([]byte("Adding products to the database...."))
}

// Get all products handler
func getProducts(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Getting all products from the database"))
}

// Get a single product hanler
func getProduct(w http.ResponseWriter, r *http.Request) {

	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil || id < 1 {
		http.NotFound(w, r)
		return
	}
	fmt.Fprintf(w, "Display a specific product with ID %d...", id)
}

func main() {

	mux := http.NewServeMux()

	mux.HandleFunc("/", dasboard)
	mux.HandleFunc("/product/create", addProduct)
	mux.HandleFunc("/product", getProducts)
	mux.HandleFunc("/product?id=1", getProduct)

	log.Println("starting a server on port 4000...")
	err := http.ListenAndServe(":4000", mux)
	log.Fatal(err)
}
