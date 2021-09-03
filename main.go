package main

import (
	"log"
	"net/http"
)

func dasboard(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Welcome to my dashboard..."))
}

// Add products handler
func addProduct(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Adding products to the database...."))
}

// Get all products handler
func getProducts(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Getting all products from the database"))
}

// Get a single product hanler
func getProduct(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Getting product with id ID..."))
}

func main() {

	mux := http.NewServeMux()

	mux.HandleFunc("/dashboard", dasboard)
	mux.HandleFunc("/product/create", addProduct)
	mux.HandleFunc("/product", getProducts)
	mux.HandleFunc("/product/id", getProduct)

	log.Println("starting a server on port 4000...")
	err := http.ListenAndServe(":4000", mux)
	log.Fatal(err)
}
