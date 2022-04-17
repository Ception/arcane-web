package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", handlerFunc)

	http.ListenAndServe(":80", nil)
}

func handlerFunc(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>Placeholder for template to come! Page Requested :: %s</h1>\n", r.URL.Path)
}
