package main

import (
	"fmt"
	"html"
	"log"
	"net/http"
)

// net/http: This package provides HTTP functionality, like creating a server and handling HTTP requests.

func main() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "hello there %q", html.EscapeString(r.URL.Path))
		log.Println(r.URL.Path)

	})

	log.Println("server started at 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))

	http.ListenAndServe("8080", nil)
}
