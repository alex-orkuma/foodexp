package main

import (
	"net/http"

	"github.com/bmizerany/pat"
	"github.com/justinas/alice"
)

func (app *application) routes() http.Handler {

	standardMiddleware := alice.New(app.recoverPanic, app.logRequest, secureHeaders)

	// Dynamic Chian middleware for dynmic routes
	dynmicMiddleware := alice.New(app.session.Enable)
	mux := pat.New()

	mux.Get("/", dynmicMiddleware.ThenFunc(app.dashboard))

	// Products routes
	mux.Get("/product/create", dynmicMiddleware.ThenFunc(app.createProductForm))
	mux.Post("/product/create", dynmicMiddleware.ThenFunc(app.addProduct))
	mux.Get("/product/:id", dynmicMiddleware.ThenFunc(app.getProduct))
	//mux.HandleFunc("/product", app.getProducts)

	// User routes
	mux.Get("/user/signup", dynmicMiddleware.ThenFunc(app.signupUserForm))
	mux.Post("/user/signup", dynmicMiddleware.ThenFunc(app.signupUser))
	mux.Get("/user/login", dynmicMiddleware.ThenFunc(app.loginUserForm))
	mux.Post("/user/login", dynmicMiddleware.ThenFunc(app.loginUser))
	mux.Post("/user/logout", dynmicMiddleware.ThenFunc(app.logoutUser))

	fileserver := http.FileServer(http.Dir("./ui/static"))

	mux.Get("/static/", http.StripPrefix("/static", fileserver))

	return standardMiddleware.Then(mux)
}
