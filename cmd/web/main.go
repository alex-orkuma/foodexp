package main

import (
	"log"
	"net/http"
)

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
