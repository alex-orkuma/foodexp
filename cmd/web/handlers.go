package main

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/alex-orkuma/foodexp/pkg/forms"
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

	// Validation check
	form := forms.New(r.PostForm)
	form.Required("food_id", "food_name", "shelf_life")
	form.MaxLength("food_id", 100)
	form.MaxLength("food_name", 100)

	if !form.Valid() {
		app.render(w, r, "createProducts.page.tmpl", &templateData{Form: form})
		return
	}

	id, err := app.products.Insert(form.Get("food_id"), form.Get("food_name"), form.Get("shelf_life"))
	if err != nil {
		app.serveError(w, err)
		return
	}

	app.session.Put(r, "flash", "Product successfully created")
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
	app.render(w, r, "createProducts.page.tmpl", &templateData{
		Form: forms.New(nil),
	})

}

func (app *application) signupUserForm(w http.ResponseWriter, r *http.Request) {
	app.render(w, r, "signup.page.tmpl", &templateData{
		Form: forms.New(nil),
	})
}

func (app *application) signupUser(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		app.clientError(w, http.StatusBadRequest)
	}

	form := forms.New(r.PostForm)
	form.Required("first_name", "last_name", "email", "password")
	form.MaxLength("first_name", 100)
	form.MaxLength("last_name", 100)
	form.MaxLength("email", 255)
	form.MatchesPattern("email", forms.EmailRX)
	form.MinLength("password", 10)

	if !form.Valid() {
		app.render(w, r, "signup.page.tmpl", &templateData{Form: form})
	}

	err = app.users.Insert(form.Get("first_name"), form.Get("last_name"), form.Get("email"), form.Get("password"), form.Get("role"))
	if err != nil {
		if errors.Is(err, models.ErrDuplicateEmail) {
			form.Errors.Add("email", "Address is already in use")
			app.render(w, r, "signup.page.tmpl", &templateData{Form: form})
		} else {
			app.serveError(w, err)
		}
		return
	}

	http.Redirect(w, r, "/user/login", http.StatusSeeOther)
}

func (app *application) loginUserForm(w http.ResponseWriter, r *http.Request) {
	app.render(w, r, "login.page.tmpl", &templateData{Form: forms.New(nil)})
}

func (app *application) loginUser(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Authenticate and login user...")
}

func (app *application) logoutUser(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Logout the user...")
}
