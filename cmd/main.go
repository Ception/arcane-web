package main

import (
	"goarcane/postgresql"
	"goarcane/web"
	"log"
	"net/http"
)

func main() {
	psqlInfo := "user=postgres password=root dbname=arcane-web sslmode=disable"

	store, err := postgresql.NewAccountStore(psqlInfo)
	if err != nil {
		log.Fatal(err)
	}

	h := web.NewHandler(store)
	http.ListenAndServe(":80", h)
}
