package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func (app *Application) serverError(err error, w http.ResponseWriter) {
	fmt.Println(err)

	http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
}

/*func gameState(idx int, theGame []displayedText) displayedText {
	return theGame[idx]
}

func printStory(stories []string) {
	for _, story := range stories {
		fmt.Println(story)
	}
	fmt.Println()
}

func choosingOption(options []option) string {
	for idx, opt := range options {
		fmt.Println(idx, ") ", opt.text)
	}
	fmt.Println("which one would you choose : ")
	var res int
	fmt.Scan(&res)
	return options[res].arc
}

func game(g displayedText) {
	fmt.Println(g.title)
	printStory(g.story)
}
*/
func render (specifiedData displayedText, w http.ResponseWriter) error{
	t, err := template.New("base").ParseFiles("./ui/html/base.page.tmpl")
	if err != nil {
		return err
	}
	err = t.Execute(w, specifiedData)
	if err != nil{
		return err
	}
	return nil
}

func (app *Application) renderSpecificPage(w http.ResponseWriter, story string) {
	idx := app.MapStoryToIdx[story]
	templateData := app.Data[idx]

	err := render(templateData, w)
	if err != nil {
		app.serverError(err, w)
	}
}