package main

import (
	"fmt"
	"html/template"
	"net/http"
	"strconv"
)

func (app *application) dashboard(w http.ResponseWriter, r *http.Request) {

	files := []string{
		"./ui/html/dashboard.page.tmpl",
		"./ui/html/sidebar.partial.tmpl",
		"./ui/html/base.layout.tmpl",
	}

	ts, err := template.ParseFiles(files...)
	if err != nil {
		app.serveError(w, err )
		return
	}

	err = ts.Execute(w, nil)
	if err != nil {
		app.serveError(w, err)
	}
}

// Add products handler
func (app *application) addProduct(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodPost {

		w.Header().Set("Allow", http.MethodPost)
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}
	w.Write([]byte("Adding products to the database...."))
}

// Get all products handler
func (app *application) getProducts(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Getting all products from the database"))
}

// Get a single product hanler
func (app *application) getProduct(w http.ResponseWriter, r *http.Request) {

	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil || id < 1 {
		app.NotFound(w)
		return
	}
	fmt.Fprintf(w, "Display a specific product with ID %d...", id)
}
