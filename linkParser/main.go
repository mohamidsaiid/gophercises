package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"golang.org/x/net/html"
)

type Link struct {
	href string
	text string
}

func main() {
	fileName := flag.String("file", "ex1.html", "get the needed html file")
	flag.Parse()
	file, err := os.Open(*fileName)
	defer file.Close()
	if err != nil {
		log.Fatal(err)
	}

	doc, err := html.Parse(file)
	if err != nil {
		log.Fatal("couldn't parse the file")
	}

	var f func(*html.Node)
	f = func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "a" {
			fmt.Print(n, " ")
			for _, a := range n.Attr {
				if a.Key == "href" {
					fmt.Println(a.Val)
					break
				}
			}
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			f(c)
		}
	}
	f(doc)
}


