package main

import "net/http"

func (app *application) routes() *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("/", app.dashboard)
	mux.HandleFunc("/product/create", app.addProduct)
	//mux.HandleFunc("/product", app.getProducts)
	mux.HandleFunc("/product", app.getProduct)

	fileserver := http.FileServer(http.Dir("./ui/static"))

	mux.Handle("/static/", http.StripPrefix("/static", fileserver))

	return mux
}
