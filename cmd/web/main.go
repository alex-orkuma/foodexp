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

	fileserver := http.FileServer(http.Dir("./ui/static"))

	mux.Handle("/static/", http.StripPrefix("/static", fileserver))

	log.Println("starting a server on port 5000...")
	err := http.ListenAndServe(":5000", mux)
	log.Fatal(err)
}
