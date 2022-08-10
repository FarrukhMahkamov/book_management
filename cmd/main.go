package main

import (
	"log"
	"net/http"

	"github.com/FarrukhMahkamov/book_management/pkg/route"
	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	route.RegisterBookRoutes(r)

	http.Handle("/", r)
	log.Fatal(http.ListenAndServe("localhost:9080", r))
}
