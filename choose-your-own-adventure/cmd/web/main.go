package main

import (
	"fmt"
	"net/http"
)

type option struct {
	Text string
	Arc  string
}

type displayedText struct {
	Arc     string
	Title   string
	Story   []string
	Options []option
}

type Application struct {
	Data          []displayedText
	MapStoryToIdx map[string]int
}

func main() {

    mpStoryToIdx, data := parsingJson()

	app := &Application{
		Data : data, 
		MapStoryToIdx: mpStoryToIdx,
	}

	serv := &http.Server{
		Addr:    ":4000",
		Handler: app.routes(),
	}

	fmt.Println("starting server on port :4000")

	err := serv.ListenAndServe()
	if err != nil {
		fmt.Println(err)
	}
	/*	state := mpStoryToIdx["intro"]

		for g := gameState(state, data); len(g.options) != 0; {
			game(g)
			res := choosingOption(g.options)
			g = gameState(mpStoryToIdx[res], data)

			if len(g.options) == 0 {
				game(g)
			}
			fmt.Println()
		}*/
}
