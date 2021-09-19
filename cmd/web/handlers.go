package main

import (
	"errors"
	"fmt"
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

	app.render(w, r, "dashboard.page.tmpl", &templateData{
		Products: p,
	})
}

// Add products handler
func (app *application) addProduct(w http.ResponseWriter, r *http.Request) {

	err := r.ParseForm()
	if err != nil {
		app.clientError(w, http.StatusBadRequest)
		return
	}

	food_id := r.PostForm.Get("food_id")
	food_name := r.PostForm.Get("food_name")
	shelf_life := r.PostForm.Get("shelf_life")

	id, err := app.products.Insert(food_id, food_name, shelf_life)

	if err != nil {
		app.serveError(w, err)
		return
	}
	http.Redirect(w, r, fmt.Sprintf("/product/%d", id), http.StatusSeeOther)
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

	id, err := strconv.Atoi(r.URL.Query().Get(":id"))
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

func (app *application) createProductForm(w http.ResponseWriter, r *http.Request) {
	app.render(w, r, "createProducts.page.tmpl", nil)

}
