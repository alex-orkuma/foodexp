package main

import (
	"errors"
	"fmt"
	"html/template"
	"net/http"
	"strconv"

	"github.com/alex-orkuma/foodexp/pkg/models"
)

func (app *application) dashboard(w http.ResponseWriter, r *http.Request) {

	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	p, err := app.products.Latest()
	if err != nil {
		app.serveError(w, err)
		return
	}

	data := &templateData{Products: p}

	files := []string{
		"./ui/html/dashboard.page.tmpl",
		"./ui/html/sidebar.partial.tmpl",
		"./ui/html/base.layout.tmpl",
	}

	ts, err := template.ParseFiles(files...)
	if err != nil {
		app.serveError(w, err)
		return
	}

	err = ts.Execute(w, data)
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

	food_id := "FI9338"
	food_name := "Akpu"
	shelf_life := "50"

	id, err := app.products.Insert(food_id, food_name, shelf_life)

	if err != nil {
		app.serveError(w, err)
		return
	}
	http.Redirect(w, r, fmt.Sprintf("/product?id=%d", id), http.StatusSeeOther)
}

// Get all products handler
func (app *application) getProducts(w http.ResponseWriter, r *http.Request) {

	p, err := app.products.Latest()
	if err != nil {
		app.serveError(w, err)
		return
	}

	for _, product := range p {
		fmt.Fprintf(w, "%v\n", product)
	}
}

// Get a single product hanler
func (app *application) getProduct(w http.ResponseWriter, r *http.Request) {

	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil || id < 1 {
		app.NotFound(w)
		return
	}

	s, err := app.products.Get(id)
	if err != nil {
		if errors.Is(err, models.ErrNoRecord) {
			app.NotFound(w)
		} else {
			app.serveError(w, err)
		}
		return
	}
	fmt.Fprintf(w, "%v", s)
}
