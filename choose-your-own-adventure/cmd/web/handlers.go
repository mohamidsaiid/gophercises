package main

import (
	"net/http"
)

func (app *Application) intro(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	app.renderSpecificPage(w, "intro")	
}

func (app *Application) newYork(w http.ResponseWriter, r *http.Request) {
	app.renderSpecificPage(w, "new-york")	
}

func (app *Application) debate(w http.ResponseWriter, r *http.Request) {
	app.renderSpecificPage(w, "debate")	
}

func (app *Application) seanKelly(w http.ResponseWriter, r *http.Request) {

	app.renderSpecificPage(w, "sean-kelly")	
}

func (app *Application) markBates(w http.ResponseWriter, r *http.Request) {

	app.renderSpecificPage(w, "mark-bates")	

}

func (app *Application) denver(w http.ResponseWriter, r *http.Request) {
	app.renderSpecificPage(w, "denver")	
}
func (app *Application) home(w http.ResponseWriter, r *http.Request) {

	app.renderSpecificPage(w, "home")	

}
