package main

import (
	"net/http"
)

func home (w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>welcome</h1>"))
}

func hello (w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>hello page</h1>"))
}
func main() {
	mux := http.NewServeMux()

	mp := map[string]http.HandlerFunc{
		"/" : home,
		"/hello" : hello,
	}
	for key, val := range mp {
		mux.HandleFunc(key, val)
	}
	http.ListenAndServe(":3000", mux)
}