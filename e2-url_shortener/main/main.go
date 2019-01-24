package main

import (
	"fmt"
	"github.com/gophercises/e2-url_shortener"
	"net/http"
)

func main () {
	mux := defaultMux()

	pathsToUrls := map[string]string {
		"/urlshort-godoc": "https://godoc.org/github.com/gophercises/urlshort",
		"/yaml-godoc": "https://godoc.org/gopkg.in/yaml.v2",
	}

	mapHandler := urlshort.MapHandler(pathsToUrls, mux)

	yaml := `
- path: /urlshort
  url: https://github.com/gophercises/urlshort
- path: /urlshort-final
  url: https://github.com/gophercises/urlshort/tree/final
`
	yamlHandler, err := urlshort.YAMLHandler([]byte(yaml), mapHandler)
	if err != nil {
		panic(err)
	}
	fmt.Println("starting the server on: 8080")
	http.ListenAndServe(":8080",yamlHandler)
	//http.ListenAndServe(":8080",mapHandler)
}

func defaultMux () *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", hello)
	return mux
}

func hello (w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "hello world!")
}