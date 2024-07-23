package main

import "net/http"

func (app *Application) routes() http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("/", app.intro)
	mux.HandleFunc("/debate", app.debate)
	mux.HandleFunc("/mark-bates", app.markBates)
	mux.HandleFunc("/new-york", app.newYork)
	mux.HandleFunc("/sean-kelly", app.seanKelly)
	mux.HandleFunc("/home", app.home)
	mux.HandleFunc("/denver", app.denver)

	return mux
}
