package main

import (
	"text/template"
	"os"
)

func main() {
	t1, err := template.Parse("value is {{.}}\n")
	if err != nil {
		panic(err)
	}
	t1.Execute(os.Stdout, 5)
	t1.Execute(os.Stdout, []string{
		"hello",
		"world",
		"golang",
	})
}