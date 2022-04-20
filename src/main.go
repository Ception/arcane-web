package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", indexHandler)

	fmt.Println("Starting webserver...")
	go http.ListenAndServe(":80", nil)
	fmt.Println("Listening on port 80...")

}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>Placeholder for template to come!</h1>")
}
