package main

import (
	"net/http"

	"github.com/bmizerany/pat"
	"github.com/justinas/alice"
)

func (app *application) routes() http.Handler {

	standardMiddleware := alice.New(app.recoverPanic, app.logRequest, secureHeaders)
	mux := pat.New()

	mux.Get("/", http.HandlerFunc(app.dashboard))
	mux.Get("/product/create", http.HandlerFunc(app.createProductForm))
	mux.Post("/product/create", http.HandlerFunc(app.addProduct))
	mux.Get("/product/:id", http.HandlerFunc(app.getProduct))
	//mux.HandleFunc("/product", app.getProducts)

	fileserver := http.FileServer(http.Dir("./ui/static"))

	mux.Get("/static/", http.StripPrefix("/static", fileserver))

	return standardMiddleware.Then(mux)
}
